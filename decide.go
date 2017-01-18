package decide

import (
	"github.com/pkg/errors"
	exp "github.com/sjhitchner/go-decide/expression"
	"github.com/sjhitchner/go-decide/lexer"
	"github.com/sjhitchner/go-decide/parser"
	"io"
)

type Logger interface {
	Appendf(f string, a ...interface{})
}

type Tree struct {
	root *Node
}

func NewTree(objects map[string][]string) (*Tree, error) {

	objectSorter := NewFrequencySorter()
	expressionSorter := NewFrequencySorter()

	// Build up frequency tables
	for object, list := range objects {
		for _, expString := range list {
			expressionSorter.AddToFrequencies(expString)
		}
		objectSorter.AddValue(object, len(list))
	}

	var root *Node
	var err error
	for _, object := range objectSorter.FrequencyList() {
		expressions := objects[object]
		expressionSorter.SortReverse(expressions)

		root, err = addNode(root, expressions, object)
		if err != nil {
			return nil, err
		}
	}

	return &Tree{root}, nil
}

func addNode(node *Node, expressions []string, object string) (*Node, error) {
	var expression exp.Expression
	var sliced []string
	var err error

	if node == nil {
		expression, sliced, err = nextExpression(expressions, nil)
		if err != nil {
			return nil, err
		}
		node = NewNode(expression)

	} else {
		expression, sliced, err = nextExpression(expressions, node.Expression)
		if err != nil {
			return nil, err
		}
	}

	if expression != nil { // There was a match
		if len(sliced) > 0 {
			node.True, err = addNode(node.True, sliced, object)
			return node, err
		}

		node.Payload = append((*node).Payload, object)
		return node, nil
	}

	// If there is no matching expression, but the expression is a negative (NOT) expression
	// the expressions are independent and the asset should be present whether or this check
	// is true or false
	// Ex: no matching app.Id != "A" expression, but the request could match the expression and still be valid for the rest of the expressions
	switch nodeExp := node.Expression.(type) {
	case *exp.ComparisonExpression:
		if nodeExp.Comparison == exp.NotEquals || nodeExp.Comparison == exp.IsNot {
			node.True, err = addNode(node.True, sliced, object)
		}
	case *exp.NegationExpression:
		node.True, err = addNode(node.True, sliced, object)
	}

	if err != nil {
		return node, err
	}

	node.False, err = addNode(node.False, expressions, object)
	return node, err
}

func nextExpression(expressions []string, match exp.Expression) (exp.Expression, []string, error) {
	if len(expressions) == 0 {
		return nil, nil, nil
	}

	if match == nil {
		// No match, use first
		expression, err := NewExpression(expressions[0])
		if err != nil {
			return nil, nil, err
		}
		return expression, expressions[1:], nil
	}

	// See if the expression exists in the list
	e := make([]string, 0, len(expressions))
	for i, exprstr := range expressions {
		expression, err := NewExpression(exprstr)
		if err != nil {
			return nil, nil, err
		}

		if expression.String() == match.String() {
			return expression, append(e, expressions[i+1:]...), nil
		}
		e = append(e, exprstr)
	}

	return nil, expressions, nil
}

func (t Tree) Evaluate(ctx exp.Context, logger Logger) ([]string, error) {
	payloadMap := make(map[string]struct{})
	if err := t.root.Evaluate(ctx, payloadMap, logger); err != nil {
		return nil, errors.Wrap(err, "Error evaluating tree")
	}

	list := make([]string, 0, len(payloadMap))
	for key := range payloadMap {
		list = append(list, key)
	}

	return list, nil
}

func (t Tree) String() string {
	return t.root.String()
}

func (t Tree) Graph(w io.Writer) error {
	return Graph(w, t.root)
}

func NewExpression(str string) (exp.Expression, error) {

	lex := lexer.NewLexer([]byte(str))
	p := parser.NewParser()
	expression, err := p.Parse(lex)
	if err != nil {
		return nil, errors.Wrapf(err, "Unable to parse expression `%s`", str)
	}

	return expression.(exp.Expression), nil
}
