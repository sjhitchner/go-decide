package decide

import (
	"fmt"
	"github.com/pkg/errors"
	exp "github.com/sjhitchner/go-decide/expression"
	"log"
)

// Generate syntax parser
//go:generate gocc -a grammar.bnf

// Decision Tree Node
type Node struct {
	Expression exp.Expression
	Payload    []string
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

func (t Node) Evaluate(ctx exp.Context, payloadMap map[string]struct{}) error {
	result, err := toBool(t.Expression.Evaluate(ctx))
	if err != nil {
		return errors.Wrapf(err, "failed to evaluate expression %v", t.Expression)
	}

	if result {
		for _, payload := range t.Payload {
			if _, ok := payloadMap[payload]; !ok {
				payloadMap[payload] = struct{}{}
			}
		}
	}

	log.Printf("EVAL %s Result: %v Payload: %v\n", t.Expression, result, payloadMap)

	if result && t.True != nil {
		return t.True.Evaluate(ctx, payloadMap)
	} else if t.False != nil {
		return t.False.Evaluate(ctx, payloadMap)
	} else if result {
		for _, payload := range t.Payload {
			if _, ok := payloadMap[payload]; !ok {
				payloadMap[payload] = struct{}{}
			}
		}
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
