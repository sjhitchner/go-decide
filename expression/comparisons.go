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

func OperationGreaterThan(ai, bi interface{}) (bool, error) {
	return ComparisonOperation(ai, bi,
		func(a, b bool) bool {
			return false
		},
		func(a, b int64) bool {
			return a > b
		},
		func(a, b float64) bool {
			return a > b
		},
		func(a, b string) bool {
			return a > b
		},
	)
}

func OperationGreaterThanEquals(ai, bi interface{}) (bool, error) {
	return ComparisonOperation(ai, bi,
		func(a, b bool) bool {
			return false
		},
		func(a, b int64) bool {
			return a >= b
		},
		func(a, b float64) bool {
			return a >= b
		},
		func(a, b string) bool {
			return a >= b
		},
	)
}

func OperationLessThan(ai, bi interface{}) (bool, error) {
	return ComparisonOperation(ai, bi,
		func(a, b bool) bool {
			return false
		},
		func(a, b int64) bool {
			return a < b
		},
		func(a, b float64) bool {
			return a < b
		},
		func(a, b string) bool {
			return a < b
		},
	)
}

func OperationLessThanEquals(ai, bi interface{}) (bool, error) {
	return ComparisonOperation(ai, bi,
		func(a, b bool) bool {
			return false
		},
		func(a, b int64) bool {
			return a <= b
		},
		func(a, b float64) bool {
			return a <= b
		},
		func(a, b string) bool {
			return a <= b
		},
	)
}

func OperationEquals(ai, bi interface{}) (bool, error) {
	return ComparisonOperation(ai, bi,
		func(a, b bool) bool {
			return a == b
		},
		func(a, b int64) bool {
			return a == b
		},
		func(a, b float64) bool {
			return a == b
		},
		func(a, b string) bool {
			return a == b
		},
	)
}

func OperationNotEquals(ai, bi interface{}) (bool, error) {
	return ComparisonOperation(ai, bi,
		func(a, b bool) bool {
			return a != b
		},
		func(a, b int64) bool {
			return a != b
		},
		func(a, b float64) bool {
			return a != b
		},
		func(a, b string) bool {
			return a != b
		},
	)
}

func OperationIs(ai interface{}, bi interface{}) (bool, error) {
	return ComparisonOperation(ai, bi,
		func(a, b bool) bool {
			return a == b
		},
		func(a, b int64) bool {
			return a == b
		},
		func(a, b float64) bool {
			return a == b
		},
		func(a, b string) bool {
			return a == b
		},
	)
}

func ComparisonOperation(
	ai, bi interface{},
	fbool func(a, b bool) bool,
	fint func(a, b int64) bool,
	ffloat func(a, b float64) bool,
	fstr func(a, b string) bool) (bool, error) {

	if ai != nil {
		switch a := ai.(type) {
		case bool:
			b, ok := bi.(bool)
			if !ok {
				if bi != nil {
					return false, errors.Wrapf(IncompatibleTypeError, "(bool) %v incompatible with %v", ai, bi)
				}
				b = false
			}
			return fbool(a, b), nil

		case int:
			b, ok := bi.(int)
			if !ok {
				if bi != nil {
					return false, errors.Wrapf(IncompatibleTypeError, "(int) %v incompatible with %v", ai, bi)
				}
				b = 0
			}
			return fint(int64(a), int64(b)), nil

		case int64:
			b, ok := bi.(int64)
			if !ok {
				if bi != nil {
					return false, errors.Wrapf(IncompatibleTypeError, "(int64) %v incompatible with %v", ai, bi)
				}
				b = 0
			}
			return fint(a, b), nil

		case float64:
			b, ok := bi.(float64)
			if !ok {
				if bi != nil {
					return false, errors.Wrapf(IncompatibleTypeError, "(float64) %v incompatible with %v", ai, bi)
				}
				b = 0
			}
			return ffloat(a, b), nil

		case string:
			b, ok := bi.(string)
			if !ok {
				if bi != nil {
					return false, errors.Wrapf(IncompatibleTypeError, "(string) %v incompatible with %v", ai, bi)
				}
				b = ""
			}
			return fstr(a, b), nil

		default:
			//return fn(a, nil)
			return false, errors.Errorf("Invalid type %v, %v", reflect.TypeOf(ai), ai)
		}

	} else if bi != nil {
		switch b := bi.(type) {
		case bool:
			a, ok := ai.(bool)
			if !ok {
				if ai != nil {
					return false, errors.Wrapf(IncompatibleTypeError, "(bool) %v incompatible with %v", ai, bi)
				}
				a = false
			}
			return fbool(a, b), nil

		case int:
			a, ok := ai.(int)
			if !ok {
				if ai != nil {
					return false, errors.Wrapf(IncompatibleTypeError, "(int) %v incompatible with %v", ai, bi)
				}
				a = 0
			}
			return fint(int64(a), int64(b)), nil

		case int64:
			a, ok := ai.(int64)
			if !ok {
				if ai != nil {
					return false, errors.Wrapf(IncompatibleTypeError, "(int64) %v incompatible with %v", ai, bi)
				}
				a = 0
			}
			return fint(a, b), nil

		case float64:
			a, ok := ai.(float64)
			if !ok {
				if ai != nil {
					return false, errors.Wrapf(IncompatibleTypeError, "(float64) %v incompatible with %v", ai, bi)
				}
				a = 0
			}
			return ffloat(a, b), nil

		case string:
			a, ok := ai.(string)
			if !ok {
				if ai != nil {
					return false, errors.Wrapf(IncompatibleTypeError, "(string) %v incompatible with %v", ai, bi)
				}
				a = ""
			}
			return fstr(a, b), nil

		default:
			//return fn(nil, b), nil
			return false, errors.Errorf("Invalid type %v, %v", reflect.TypeOf(ai), ai)
		}
	}

	return false, nil
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
