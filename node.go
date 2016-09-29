package decide

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sjhitchner/go-decide/expression"
)

// Generate syntax parser
//go:generate gocc -a grammar.bnf

type Object string

type EvaluationLog interface {
	Add(objects ...Object)
}

type evaluationLog struct {
	objects []Object
}

func NewEvaluationLog() EvaluationLog {
	return &evaluationLog{
		make([]Object, 0, 10),
	}
}

func (t *evaluationLog) Add(objects ...Object) {
	t.objects = append(t.objects, objects...)
}

// Decision Tree Node
type Node struct {
	Expression expression.Expression
	Payload    []Object
	True       *Node
	False      *Node
}

func NewNode(expression expression.Expression) *Node {
	return &Node{
		expression,
		nil,
		nil,
		nil,
	}
}

func (t Node) Evaluate(ctx Context, log EvaluationLog) error {

	if t.Expression == nil {
		return nil
	}

	/*
		result, err := toBool(t.expression.Evaluate(ctx))
		if err != nil {
			return errors.Wrapf(err, "Failed to evaluate expression %v", t.expression)
		}

		if result {
			log.Add(t.Payload...)
		}

		if result && t.True != nil {
			return t.True.Evaluate(ctx, log)
		} else if t.False != nil {
			return t.False.Evaluate(ctx, log)
		}
	*/

	return nil
}

func toBool(result interface{}, err error) (bool, error) {
	if err != nil {
		return false, err
	}

	b, ok := result.(bool)
	if !ok {
		return false, errors.Errorf("Expect bool got %v", result)
	}

	return b, nil
}

func (t Node) String() string {
	return fmt.Sprintf("(%s [%s] %s)", t.True, t.Expression, t.False)
}
