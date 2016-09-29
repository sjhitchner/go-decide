package decide

import (
	//"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	exp "github.com/sjhitchner/go-decide/expression"
)

// Generate syntax parser
//go:generate gocc -a grammar.bnf

type Object string

/*
type EvaluationLog interface {
	Add(objects ...Object)
}
*/

/*
type EvaluationLog struct {
	objects []Object `json:"objects"`
}

func NewEvaluationLog() *EvaluationLog {
	return &EvaluationLog{
		make([]Object, 0, 10),
	}
}

func (t *EvaluationLog) Add(objects ...Object) {
	t.objects = append(t.objects, objects...)
}

func (t EvaluationLog) String() string {
	b, _ := json.MarshalIndent(t, "", "  ")
	return string(b)
}
*/

// Decision Tree Node
type Node struct {
	Expression exp.Expression
	Payload    []Object
	True       *Node
	False      *Node
}

func NewNode(expression exp.Expression) *Node {
	return &Node{
		expression,
		nil,
		nil,
		nil,
	}
}

/*
	if t.Expression == nil {
		return nil
	}
*/
func (t Node) Evaluate(ctx exp.Context, log *[]Object) error {
	fmt.Println("EVAL", t.Expression, ctx) //, result, log)

	result, err := toBool(t.Expression.Evaluate(ctx))
	fmt.Println("EVAL Result:", result)
	if err != nil {
		return errors.Wrapf(err, "Failed to evaluate expression %v", t.Expression)
	}

	if result {
		//log.Add(t.Payload...)
		*log = append(*log, t.Payload...)
	}

	if result && t.True != nil {
		return t.True.Evaluate(ctx, log)
	} else if t.False != nil {
		return t.False.Evaluate(ctx, log)
	}

	return nil
}

func toBool(result interface{}, err error) (bool, error) {
	if err != nil {
		return false, err
	}

	if result == nil {
		return false, nil
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
