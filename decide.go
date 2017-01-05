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
	expression, err := NewExpression(expressions[0])
	if err != nil {
		return nil, err
	}

	if node == nil {
		node = NewNode(expression)
	}

	if node.Expression.String() == expression.String() {
		if len(expressions) > 1 {
			node.True, err = addNode(node.True, expressions[1:], object)
			return node, err
		}

		node.Payload = append((*node).Payload, object)
		return node, nil
	}

	node.False, err = addNode(node.False, expressions, object)
	return node, err
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
		return nil, errors.Wrap(err, "Unable to build AST")
	}

	return expression.(exp.Expression), nil
}
