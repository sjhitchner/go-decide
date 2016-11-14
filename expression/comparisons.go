package expression

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

var IncompatibleTypeError = errors.New("Incompatible Types Error")

// Operation

type ComparisonOperations interface {
	OperationAnd(ai interface{}, bi interface{}) (bool, error)
	OperationOr(ai interface{}, bi interface{}) (bool, error)
	OperationNot(ai interface{}) (interface{}, error)
	OperationGreaterThan(ai interface{}, bi interface{}) (bool, error)
	OperationGreaterThanEquals(ai interface{}, bi interface{}) (bool, error)
	OperationLessThan(ai interface{}, bi interface{}) (bool, error)
	OperationLessThanEquals(ai interface{}, bi interface{}) (bool, error)
	OperationEquals(ai interface{}, bi interface{}) (bool, error)
	OperationNotEquals(ai interface{}, bi interface{}) (bool, error)
	OperationIs(ai interface{}, bi interface{}) (bool, error)
	OperationContains(ai interface{}, bi interface{}) (bool, error)
}

type ReflectionComparisionOperations struct{}

func (t ReflectionComparisonOperations) OperationAnd(ai interface{}, bi interface{}) (bool, error) {
	a, oka := ai.(bool)
	b, okb := bi.(bool)

	if oka && okb {
		return a && b, nil
	}

	return false, errors.Wrapf(IncompatibleTypeError, "AND %v incompatible with %v", a, b)
}

func (t ReflectionComparisonOperations) OperationOr(ai interface{}, bi interface{}) (bool, error) {
	a, oka := ai.(bool)
	b, okb := bi.(bool)

	if oka && okb {
		return a || b, nil
	}

	return false, errors.Wrapf(IncompatibleTypeError, "OR %v incompatible with %v", a, b)
}

func (t ReflectionComparisonOperations) OperationNot(ai interface{}) (interface{}, error) {
	a, oka := ai.(bool)
	if oka {
		return !a, nil
	}

	return false, errors.Wrapf(IncompatibleTypeError, "NOT %v invalid type", a)
}

func (t ReflectionComparisonOperations) OperationGreaterThan(ai interface{}, bi interface{}) (bool, error) {
}
func (t ReflectionComparisonOperations) OperationGreaterThanEquals(ai interface{}, bi interface{}) (bool, error) {
}
func (t ReflectionComparisonOperations) OperationLessThan(ai interface{}, bi interface{}) (bool, error) {
}
func (t ReflectionComparisonOperations) OperationLessThanEquals(ai interface{}, bi interface{}) (bool, error) {
}
func (t ReflectionComparisonOperations) OperationEquals(ai interface{}, bi interface{}) (bool, error) {
}
func (t ReflectionComparisonOperations) OperationNotEquals(ai interface{}, bi interface{}) (bool, error) {
}
func (t ReflectionComparisonOperations) OperationIs(ai interface{}, bi interface{}) (bool, error) {}
func (t ReflectionComparisonOperations) OperationContains(ai interface{}, bi interface{}) (bool, error) {
}

type TypeCheckComparisonOperations struct{}

func (t TypeCheckComparisonOperations) OperationAnd(ai interface{}, bi interface{}) (bool, error) {
	a, oka := ai.(bool)
	b, okb := bi.(bool)

	if oka && okb {
		return a && b, nil
	}

	return false, errors.Wrapf(IncompatibleTypeError, "AND %v incompatible with %v", a, b)
}

func (t TypeCheckComparisonOperations) OperationOr(ai interface{}, bi interface{}) (bool, error) {
	a, oka := ai.(bool)
	b, okb := bi.(bool)

	if oka && okb {
		return a || b, nil
	}

	return false, errors.Wrapf(IncompatibleTypeError, "OR %v incompatible with %v", a, b)
}

func (t TypeCheckComparisonOperations) OperationNot(ai interface{}) (interface{}, error) {
	a, oka := ai.(bool)
	if oka {
		return !a, nil
	}

	return false, errors.Wrapf(IncompatibleTypeError, "NOT %v invalid type", a)
}

func (t TypeCheckComparisonOperations) OperationGreaterThan(ai interface{}, bi interface{}) (bool, error) {
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

func (t TypeCheckComparisonOperations) OperationGreaterThanEquals(ai interface{}, bi interface{}) (bool, error) {
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

func (t TypeCheckComparisonOperations) OperationLessThan(ai interface{}, bi interface{}) (bool, error) {
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

func (t TypeCheckComparisonOperations) OperationLessThanEquals(ai interface{}, bi interface{}) (bool, error) {
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

func (t TypeCheckComparisonOperations) OperationEquals(ai interface{}, bi interface{}) (bool, error) {
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

func (t TypeCheckComparisonOperations) OperationNotEquals(ai interface{}, bi interface{}) (bool, error) {
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

func (t TypeCheckComparisonOperations) OperationIs(ai interface{}, bi interface{}) (bool, error) {
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

func (t TypeCheckComparisonOperations) OperationContains(ai interface{}, bi interface{}) (bool, error) {
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
