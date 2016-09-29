package decide

import (
	"fmt"
	"github.com/pkg/errors"
	exp "github.com/sjhitchner/go-decide/expression"
	"github.com/sjhitchner/go-decide/lexer"
	"github.com/sjhitchner/go-decide/parser"
	"io"
)

type Context map[string]interface{}

type Tree struct {
	root *Node
}

func NewTree(objects map[Object][]string) (*Tree, error) {

	sorter := NewFrequencySorter()

	// Build up frequency table
	for _, list := range objects {
		for _, expString := range list {
			sorter.AddToFrequencies(expString)
		}
	}

	root, err := buildTree(sorter.FrequencyList())
	if err != nil {
		return nil, err
	}

	if root == nil {
		return nil, errors.New("Unable to build tree, tree is empty")
	}

	for object, list := range objects {
		sorter.Sort(list)
		if err := addObject(root, object, list); err != nil {
			return nil, err
		}
	}

	pruneTree(root)

	return &Tree{root}, nil
}

// Build tree using sorted list of most common expressions
func buildTree(expressionList []string) (*Node, error) {
	var root *Node

	if len(expressionList) == 0 {
		return nil, errors.New("Expression list is empty")
	}

	for _, str := range expressionList {
		expression, err := NewExpression(str)
		if err != nil {
			return nil, errors.Wrapf(err, "Error creating expression for %s", str)
		}
		root = addExpression(root, expression)
	}

	return root, nil
}

// Add new expression to tree
// Tree is balanced and is a deep as there are expressions
func addExpression(node *Node, expression exp.Expression) *Node {
	if node == nil {
		return NewNode(expression)
	}

	node.True = addExpression(node.True, expression)
	node.False = addExpression(node.False, expression)
	return node
}

func addObject(node *Node, object Object, expressions []string) error {

	expression, err := NewExpression(expressions[0])
	if err != nil {
		return err
	}

	if node.Expression.String() == expression.String() {

		if len(expressions) == 1 {
			node.Payload = append(node.Payload, object)
			return nil
		}

		return addObject(node.True, object, expressions[1:])

	} else {
		err := addObject(node.True, object, expressions)
		if err != nil {
			return err
		}
		return addObject(node.False, object, expressions)
	}

	return nil
}

func pruneTree(node *Node) {

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

func (t Tree) Evaluate(ctx Context) (EvaluationLog, error) {
	log := NewEvaluationLog()
	if err := t.root.Evaluate(ctx, log); err != nil {
		return nil, errors.Wrap(err, "Error evaluating tree")
	}
	return log, nil
}

func (t Tree) String() string {
	return t.root.String()
}

func (t Tree) Graph(w io.Writer) error {
	return Graph(w, t.root)
}

/*

func FindExpression(node *Node, expression expression.Expression) *Node {
	if node == nil {
		return nil
	}

	if node.Expression == expression {
		return node
	}

	result := FindExpression(node.True, expression)
	if result == nil {
		result = FindExpression(node.False, expression)
	}

	return result
}

func AddObject(tree *Node, expresions []expression.Expression, object Object) error {
	return nil
}
*/
