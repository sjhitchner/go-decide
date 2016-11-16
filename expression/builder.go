package expression

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sjhitchner/go-decide/token"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Attrib interface{}

func NewComparisonGreaterThan(a, b Attrib) (*ComparisonExpression, error) {
	return &ComparisonExpression{a.(Expression), b.(Expression), GreaterThan}, nil
}

func NewComparisonGreaterThanEquals(a, b Attrib) (*ComparisonExpression, error) {
	return &ComparisonExpression{a.(Expression), b.(Expression), GreaterThanEquals}, nil
}

func NewComparisonLessThan(a, b Attrib) (*ComparisonExpression, error) {
	return &ComparisonExpression{a.(Expression), b.(Expression), LessThan}, nil
}

func NewComparisonLessThanEquals(a, b Attrib) (*ComparisonExpression, error) {
	return &ComparisonExpression{a.(Expression), b.(Expression), LessThanEquals}, nil
}

func NewComparisonEquals(a, b Attrib) (*ComparisonExpression, error) {
	return &ComparisonExpression{a.(Expression), b.(Expression), IsEquals}, nil
}

func NewComparisonNotEquals(a, b Attrib) (*ComparisonExpression, error) {
	return &ComparisonExpression{a.(Expression), b.(Expression), NotEquals}, nil
}

func NewComparisonIsNot(a, b Attrib) (*ComparisonExpression, error) {
	fmt.Println("XXX", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
	return &ComparisonExpression{a.(Expression), b.(Expression), IsNot}, nil
}

func NewComparisonIs(a, b Attrib) (*ComparisonExpression, error) {
	return &ComparisonExpression{a.(Expression), b.(Expression), Is}, nil
}

func NewComparisonContains(a, b Attrib) (*ComparisonExpression, error) {
	return &ComparisonExpression{a.(Expression), b.(Expression), Contains}, nil
}

func NewMatches(a, b Attrib) (*RegexExpression, error) {
	blit, ok := b.(*LiteralExpression)
	if !ok {
		return nil, errors.Errorf("Regex must be a string literal not (%v)", reflect.TypeOf(b))
	}

	switch bstring := blit.Value.(type) {
	case string:
		bregex, err := regexp.Compile(bstring)
		if err != nil {
			return nil, err
		}
		return &RegexExpression{a.(Expression), bregex}, nil

	default:
		return nil, errors.Errorf("Literal value not a string (%v) %v", bstring, reflect.TypeOf(bstring))
	}
}

func NewLiteralBool(a Attrib) (*LiteralExpression, error) {
	abool, ok := a.(bool)
	if !ok {
		return nil, errors.Errorf("%v not a bool", a)
	}
	return &LiteralExpression{abool}, nil
}

func NewLiteralInt(a Attrib) (*LiteralExpression, error) {
	aint, err := IntValue(a)
	if err != nil {
		return nil, err
	}
	return &LiteralExpression{aint}, nil
}

func NewLiteralFloat(a Attrib) (*LiteralExpression, error) {
	afloat, err := FloatValue(a)
	if err != nil {
		return nil, err
	}
	return &LiteralExpression{afloat}, nil
}

func NewLiteralString(a Attrib) (*LiteralExpression, error) {
	astring := StringValue(a)
	astring = strings.TrimPrefix(astring, `'`)
	astring = strings.TrimPrefix(astring, `"`)
	astring = strings.TrimSuffix(astring, `'`)
	astring = strings.TrimSuffix(astring, `"`)
	return &LiteralExpression{astring}, nil
}

func NewLogicalOr(a, b Attrib) (*LogicalExpression, error) {
	return &LogicalExpression{a.(Expression), b.(Expression), Or}, nil
}

func NewLogicalAnd(a, b Attrib) (*LogicalExpression, error) {
	return &LogicalExpression{a.(Expression), b.(Expression), And}, nil
}

func NewNegation(a Attrib) (*NegationExpression, error) {
	return &NegationExpression{a.(Expression)}, nil
}

func NewResolver(a Attrib) (*ResolverExpression, error) {
	key := StringValue(a)
	return &ResolverExpression{key}, nil
}

func NewNull() (*LiteralExpression, error) {
	return &LiteralExpression{nil}, nil
}

func StringValue(ai Attrib) string {
	switch a := ai.(type) {
	case *token.Token:
		return string(a.Lit)
	case string:
		return a
	default:
		return ""
	}
}

func IntValue(ai Attrib) (int64, error) {
	switch a := ai.(type) {
	case *token.Token:
		str := string(a.Lit)
		return strconv.ParseInt(str, 10, 64)
	case int64:
		return a, nil
	case int:
		return int64(a), nil
	default:
		return 0, errors.New("invalid int interface type")
	}
}

func BoolValue(ai Attrib) (bool, error) {
	switch a := ai.(type) {
	case *token.Token:
		str := string(a.Lit)
		return strconv.ParseBool(str)
	case bool:
		return a, nil
	default:
		return false, errors.New("invalid int interface type")
	}
}

func FloatValue(ai Attrib) (float64, error) {
	switch a := ai.(type) {
	case *token.Token:
		str := string(a.Lit)
		return strconv.ParseFloat(str, 64)
	case float64:
		return a, nil
	case float32:
		return float64(a), nil
	default:
		return 0.0, errors.New("invalid int interface type")
	}
}
