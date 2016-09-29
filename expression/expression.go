package expression

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"regexp"
)

const (
	And Logical = iota
	Or

	GreaterThan Comparison = iota
	GreaterThanEquals
	LessThan
	LessThanEquals
	Equals
	NotEquals
	Is
	IsNot
	Contains
)

type Context map[string]interface{}
type Logical int
type Comparison int

type Expression interface {
	Evaluate(Context) (interface{}, error)
	String() string
}

var LogicalError = errors.New("Logical Operation Error")
var ComparisonError = errors.New("Comparison Error")

// ComparisonExpression
// Contains expression that supports the following comparisons
// > >= < <= == != is contains
type ComparisonExpression struct {
	Left       Expression
	Right      Expression
	Comparison Comparison
}

func (t ComparisonExpression) Evaluate(ctx Context) (interface{}, error) {

	if t.Left == nil {
		return nil, errors.New("Left node is nil")
	}

	if t.Right == nil {
		return nil, errors.New("Right node is nil")
	}

	left, err := t.Left.Evaluate(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "Left comparison evaluate failed")
	}

	right, err := t.Right.Evaluate(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "Right comparison evaluate failed")
	}

	switch t.Comparison {
	case GreaterThan:
		return OperationGreaterThan(left, right)
	case GreaterThanEquals:
		return OperationGreaterThanEquals(left, right)
	case LessThan:
		return OperationLessThan(left, right)
	case LessThanEquals:
		return OperationLessThanEquals(left, right)
	case Equals:
		return OperationEquals(left, right)
	case NotEquals:
		return OperationNotEquals(left, right)
	case Is:
		return OperationIs(left, right)
	case Contains:
		return OperationContains(left, right)
	}

	return nil, ComparisonError
}

func (t ComparisonExpression) String() string {
	var l, r string

	if t.Left != nil {
		l = t.Left.String()
	} else {
		l = "empty"
	}

	if t.Right != nil {
		r = t.Right.String()
	} else {
		r = "empty"
	}

	return fmt.Sprintf("(%s %s %s)", l, t.Comparison, r)
}

/* Logical Node

type LogicalNode struct {
	Left    Node
	Right   Node
	Logical Logical
	Objects []Object
}

func (t LogicalNode) Evaluate(ctx Context) (interface{}, error) {
	if t.Left == nil {
		return nil, errors.New("Left node is nil")
	}
	if t.Right == nil {
		return nil, errors.New("Right node is nil")
	}

	left, err := t.Left.Evaluate(tracer, ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "Left logical evaluate failed")
	}

	right, err := t.Right.Evaluate(tracer, ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "Right logical evaluate failed")
	}

	var result bool
	switch t.Logical {
	case And:
		result, err = OperationAnd(left, right)
	case Or:
		result, err = OperationOr(left, right)
	default:
		result, err = false, LogicalError
	}

	return evaluateResult(tracer, result, err, t.Objects)
}

func (t LogicalNode) String() string {
	return nodeToString(t.Left, t.Right, t.Logical)
}
*/

// Regex Expression

type RegexExpression struct {
	Node  Expression
	regex *regexp.Regexp
}

func (t RegexExpression) Evaluate(ctx Context) (interface{}, error) {
	node, err := t.Node.Evaluate(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "Regex evaluating node failed")
	}

	str, ok := node.(string)
	if !ok {
		return false, errors.Wrapf(IncompatibleTypeError, "REGEX %v incompatible with %v", str, reflect.TypeOf(str))
	}

	return t.regex.MatchString(node.(string)), nil
}

func (t RegexExpression) String() string {
	return fmt.Sprintf("%s /%s/", t.Node.String(), t.regex.String())
}

// ClauseExpression

type ClauseExpression struct {
	Expression Expression
}

func (t ClauseExpression) Evaluate(ctx Context) (interface{}, error) {
	return t.Expression.Evaluate(ctx)
}

func (t ClauseExpression) String() string {
	return fmt.Sprintf("(%s)", t.Expression)
}

// NegateExpression

type NegationExpression struct {
	Expression Expression
}

func (t NegationExpression) Evaluate(ctx Context) (interface{}, error) {
	a, err := t.Expression.Evaluate(ctx)
	if err != nil {
		return nil, nil
	}
	return OperationNot(a)
}

func (t NegationExpression) String() string {
	return fmt.Sprintf("NOT %s", t.Expression)
}

// LiteralExpression

type LiteralExpression struct {
	Value interface{}
}

func (t LiteralExpression) Evaluate(ctx Context) (interface{}, error) {
	return t.Value, nil
}

func (t LiteralExpression) String() string {
	return fmt.Sprintf("'%v'", t.Value)
}

// ResolverExpression

type ResolverExpression struct {
	key string
}

func (t ResolverExpression) Evaluate(ctx Context) (interface{}, error) {
	// TODO arbitrarily deep context maps
	ai, ok := ctx[t.key]
	if !ok {
		// TODO need to return nil is key doesn't exist
		// return false, errors.Wrapf(ContextMissingKeyError, "key %s doesn't exist", t.key)
		return nil, nil
	}

	switch a := ai.(type) {
	case int:
		return int64(a), nil
	case float32:
		return float64(a), nil
	default:
		return a, nil
	}
}

func (t ResolverExpression) String() string {
	return fmt.Sprintf("$%s", t.key)
}

func (t Logical) String() string {
	switch t {
	case And:
		return "AND"
	case Or:
		return "OR"
	default:
		return "unknown"
	}
}

func (t Comparison) String() string {
	switch t {
	case GreaterThan:
		return ">"
	case GreaterThanEquals:
		return ">="
	case LessThan:
		return "<"
	case LessThanEquals:
		return "<="
	case Equals:
		return "=="
	case NotEquals:
		return "!="
	case Is:
		return "is"
	case Contains:
		return "contains"
	default:
		return "unknown"
	}
}
