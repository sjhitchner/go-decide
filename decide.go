package decide

import (
	"github.com/pkg/errors"
	exp "github.com/sjhitchner/go-decide/expression"
	"github.com/sjhitchner/go-decide/lexer"
	"github.com/sjhitchner/go-decide/parser"
	"io"
)

type Tree struct {
	root *Node
}

func NewTree(objects map[string][]string) (*Tree, error) {

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
		if err := addPayload(root, object, list); err != nil {
			return nil, err
		}
	}

	pruneTree(root)

	return &Tree{root}, nil
}

func (t Tree) Evaluate(ctx exp.Context) ([]string, error) {
	var list []string
	if err := t.root.Evaluate(ctx, &list); err != nil {
		return nil, errors.Wrap(err, "Error evaluating tree")
	}
	return list, nil
}

func (t Tree) String() string {
	return t.root.String()
}

func (t Tree) Graph(w io.Writer) error {
	return Graph(w, t.root)
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

// addPayload
// Adding objects to the tree where they match expressions
func addPayload(node *Node, object string, expressions []string) error {

	expression, err := NewExpression(expressions[0])
	if err != nil {
		return err
	}

	if node.Expression.String() == expression.String() {

		if len(expressions) == 1 {
			node.Payload = append(node.Payload, object)
			return nil
		}

		return addPayload(node.True, object, expressions[1:])

	} else {
		err := addPayload(node.True, object, expressions)
		if err != nil {
			return err
		}
		return addPayload(node.False, object, expressions)
	}

	return nil
}

// pruneTree
// Prune nodes that have no children and nothing in the payload
func pruneTree(node *Node) bool {

	if node.True != nil {
		if prune := pruneTree(node.True); prune {
			node.True = nil
		}
	}

	if node.False != nil {
		if prune := pruneTree(node.False); prune {
			node.False = nil
		}
	}

	if node.True == nil && node.False == nil {
		// Is Leaf
		if len(node.Payload) == 0 {
			return true
		}
	}

	return false
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
