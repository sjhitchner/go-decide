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

func NewTree(objects, priority map[string][]string) (*Tree, error) {
	objectSorter := NewFrequencySorter()
	expressionSorter := NewFrequencySorter()
	// prioritySorter := NewFrequencySorter()

	// Build up frequency tables
	for object, list := range objects {
		for _, expString := range list {
			expressionSorter.AddToFrequencies(expString)
		}

		/*
			// Do a sort on the objects with priority expressions first
			if _, ok := priority[object]; ok && len(priority[object]) > 0 {
				prioritySorter.AddValue(object, len(list)+len(priority[object]))
				continue
			}
		*/

		// Do not need to sort objects that do not have priority expressions if they are already in the priority sort
		objectSorter.AddValue(object, len(list))
	}

	var root *Node
	var err error

	/*
		// Add the nodes with the priority expressions (which are sorted by length of expressions) first
		for _, object := range prioritySorter.FrequencyList() {
			expressions := objects[object]
			expressionSorter.SortReverse(expressions)

			// prepend the expressions with the priority values so they are built first
			expressions = append(priority[object], expressions...)

			root, err = addNode(root, expressions, object)
			if err != nil {
				return nil, err
			}
		}
	*/

	// Add the rest of the nodes
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

	/* Old logic
	if expression != nil { // There was a match
		if len(sliced) > 0 {
			node.True, err = addNode(node.True, sliced, object)
			return node, err
		}

		node.Payload = append((*node).Payload, object)
		return node, nil
	}

	// If there is no matching expression, but the expression is a negative (NOT) expression
	// the expressions are independent and the object should be present whether or this check
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
	*/

	if expression != nil { // There was a match
		if len(sliced) > 0 {
			if len(node.True) > 0 {
				for i, nodeTrue := range node.True {
					exp, _, _ := nextExpression(sliced, nodeTrue.Expression)
					if exp != nil { // One of the true nodes matches something left
						node.True[i], err = addNode(node.True[i], sliced, object)
						return node, err
					}
				}
			}

			// None of the True sets matched
			node.True = append(node.True, nil)
			node.True[len(node.True)-1], err = addNode(node.True[len(node.True)-1], sliced, object)
			return node, err
		}

		node.Payload = append((*node).Payload, object)
		return node, nil
	}

	// If there is no matching expression, but the expression is a negative (NOT) expression
	// the expressions are independent and the object should be present whether or this check
	// is true or false
	// Ex: no matching app.Id != "A" expression, but the request could match the expression and still be valid for the rest of the expressions
	switch nodeExp := node.Expression.(type) {
	case *exp.ComparisonExpression:
		if nodeExp.Comparison == exp.NotEquals || nodeExp.Comparison == exp.IsNot {
			node.True = append(node.True, nil)
			node.True[len(node.True)-1], err = addNode(node.True[len(node.True)-1], sliced, object)
		}
	case *exp.NegationExpression:
		node.True = append(node.True, nil)
		node.True[len(node.True)-1], err = addNode(node.True[len(node.True)-1], sliced, object)
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

func (t Tree) Evaluate(ctx exp.Context, trace Logger) ([]string, error) {
	payloadMap := make(map[string]struct{})
	if err := t.root.Evaluate(ctx, payloadMap, trace); err != nil {
		return nil, errors.Wrap(err, "Error evaluating tree")
	}

	list := make([]string, 0, len(payloadMap))
	for key := range payloadMap {
		list = append(list, key)
	}

	return list, nil
}

func (t Tree) String() string {
	if t.Empty() {
		return ""
	}
	return t.root.String()
}

func (t Tree) Graph(w io.Writer) error {
	if t.Empty() {
		return errors.New("Cannot graph empty tree")
	}
	return Graph(w, t.root)
}

func (t Tree) Empty() bool {
	return t.root == nil
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
