package expression

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

var IncompatibleTypeError = errors.New("Incompatible Types Error")

// Operation

func OperationAnd(ai interface{}, bi interface{}) (bool, error) {
	a, oka := ai.(bool)
	b, okb := bi.(bool)

	if oka && okb {
		return a && b, nil
	}

	return false, errors.Wrapf(IncompatibleTypeError, "AND %v incompatible with %v", a, b)
}

func OperationOr(ai interface{}, bi interface{}) (bool, error) {
	a, oka := ai.(bool)
	b, okb := bi.(bool)

	if oka && okb {
		return a || b, nil
	}

	return false, errors.Wrapf(IncompatibleTypeError, "OR %v incompatible with %v", a, b)
}

func OperationNot(ai interface{}) (interface{}, error) {
	a, oka := ai.(bool)
	if oka {
		return !a, nil
	}

	return false, errors.Wrapf(IncompatibleTypeError, "NOT %v invalid type", a)
}

func OperationGreaterThan(ai interface{}, bi interface{}) (bool, error) {
	result, err := ComparisonOperation(
		ai,
		bi,
		func(a, b float64) bool {
			return a > b
		},
		func(a, b string) bool {
			return a > b
		},
	)
	if err != nil {
		return false, errors.Wrap(err, "GreaterThan Error")
	}
	return result, nil
}

func OperationGreaterThanEquals(ai interface{}, bi interface{}) (bool, error) {
	result, err := ComparisonOperation(
		ai,
		bi,
		func(a, b float64) bool {
			return a >= b
		},
		func(a, b string) bool {
			return a >= b
		},
	)
	if err != nil {
		return false, errors.Wrap(err, "GreaterThanEquals Error")
	}
	return result, nil
}

func OperationLessThan(ai interface{}, bi interface{}) (bool, error) {
	result, err := ComparisonOperation(
		ai,
		bi,
		func(a, b float64) bool {
			return a < b
		},
		func(a, b string) bool {
			return a < b
		},
	)
	if err != nil {
		return false, errors.Wrap(err, "LessThan Error")
	}
	return result, nil
}

func OperationLessThanEquals(ai interface{}, bi interface{}) (bool, error) {
	result, err := ComparisonOperation(
		ai,
		bi,
		func(a, b float64) bool {
			return a <= b
		},
		func(a, b string) bool {
			return a <= b
		},
	)
	if err != nil {
		return false, errors.Wrap(err, "LessThanEquals Error")
	}
	return result, nil
}

func OperationEquals(ai interface{}, bi interface{}) (bool, error) {
	result, err := ComparisonOperation(
		ai,
		bi,
		func(a, b float64) bool {
			return a == b
		},
		func(a, b string) bool {
			return a == b
		},
	)
	if err != nil {
		return false, errors.Wrap(err, "Equals Error")
	}
	return result, nil
}

func OperationNotEquals(ai interface{}, bi interface{}) (bool, error) {
	result, err := ComparisonOperation(
		ai,
		bi,
		func(a, b float64) bool {
			return a != b
		},
		func(a, b string) bool {
			return a != b
		},
	)
	if err != nil {
		return false, errors.Wrap(err, "NotEquals Error")
	}
	return result, nil
}

func OperationIs(ai interface{}, bi interface{}) (bool, error) {
	result, err := ComparisonOperation(
		ai,
		bi,
		func(a, b float64) bool {
			return a == b
		},
		func(a, b string) bool {
			return a == b
		},
	)
	if err != nil {
		return false, errors.Wrap(err, "Is Error")
	}
	return result, nil
}

func ComparisonOperation(ai interface{}, bi interface{}, f func(a, b float64) bool, fs func(a, b string) bool) (bool, error) {
	switch a := ai.(type) {
	case int64:
		switch b := bi.(type) {
		case int64:
			return f(float64(a), float64(b)), nil
		case float64:
			return f(float64(a), b), nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "(int) %v incompatible with %v", ai, bi)

	case float64:
		switch b := bi.(type) {
		case int64:
			return f(a, float64(b)), nil
		case float64:
			return f(a, b), nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "(float) %v incompatible with %v", ai, bi)

	case string:
		b, ok := bi.(string)
		if ok {
			return fs(a, b), nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "(string) %v incompatible with %v", ai, bi)

	default:
		return false, errors.Wrapf(IncompatibleTypeError, "%v invalid type", a)
	}
}

func OperationContains(ai interface{}, bi interface{}) (bool, error) {
	switch a := ai.(type) {
	case string:
		b, ok := bi.(string)
		if ok {
			return strings.Contains(a, b), nil
		}
		return false, errors.Wrapf(IncompatibleTypeError, "contains %v incompatible with %v", ai, bi)

	case []string:
		b, ok := bi.(string)
		if !ok {
			return false, errors.Wrapf(IncompatibleTypeError, "contains %v incompatible with %v", ai, bi)
		}
		for _, ae := range a {
			if b == ae {
				return true, nil
			}
		}
		return false, nil

	case []int64:
		b, ok := bi.(int64)
		if !ok {
			return false, errors.Wrapf(IncompatibleTypeError, "contains %v incompatible with %v", ai, bi)
		}
		for _, ae := range a {
			if b == ae {
				return true, nil
			}
		}
		return false, nil

	default:
		return false, errors.Wrapf(IncompatibleTypeError, "contains %v invalid type %v", a, reflect.TypeOf(a))
	}
}
