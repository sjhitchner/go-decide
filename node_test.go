package decide

import (
	"fmt"
	exp "github.com/sjhitchner/go-decide/expression"
	. "gopkg.in/check.v1"
	"os"
	//"os/exec"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type DecisionSuite struct {
	Objects map[string][]string
	Tree    *Tree
}

var _ = Suite(&DecisionSuite{})

type TestContext map[string]interface{}

func (t TestContext) Get(key string) (interface{}, bool) {
	result, ok := t[key]
	return result, ok
}

type TestLogger struct {
	Trace []string
}

func (t *TestLogger) Appendf(f string, a ...interface{}) {
	t.Trace = append(t.Trace, fmt.Sprintf(f, a...))
}

func (s *DecisionSuite) SetUpSuite(c *C) {
	s.Objects = map[string][]string{
		"object01": []string{
			`geo_code matches '^(.*,)?(US)$'`,
		},
		"object02": []string{
			`platform = 'iOS'`,
			`device.gender != "male"`,
		},
		"object03": []string{
			`platform = 'iOS'`,
			`geo_code matches '^(.*,)?(US)$'`,
		},
		"object04": []string{
			`platform = 'Android'`,
		},
		"object05": []string{
			`platform = 'Android'`,
			`geo_code matches '^(.*,)?(US)$'`,
		},
		"object06": []string{
			`device.age_group != "13-17"`,
			`geo_code matches '^(.*,)?(US)$'`,
		},
		"object07": []string{
			`geo_code matches '^(.*,)?(US)$'`,
			`platform = 'iOS'`,
			`device.gender != "male"`,
		},
	}

	tree, err := NewTree(s.Objects)
	c.Assert(err, IsNil)

	f, err := os.Create("decision.dot")
	c.Assert(err, IsNil)
	defer f.Close()
	tree.Graph(f)

	//cmd := exec.Command("dot", "-Tpdf", "decision.dot")
	//pdf, err := os.Create("decision.pdf")
	//c.Assert(err, IsNil)
	//defer pdf.Close()
	//cmd.Stdout = pdf
	//c.Assert(cmd.Run(), IsNil)
	//fmt.Println(syscall.Exec("open decision.png", nil, nil))

	s.Tree = tree
}

func (s *DecisionSuite) Test_Context1(c *C) {
	context := TestContext{
		"geo_code":         "US",
		"platform":         "iOS",
		"device.age_group": "60",
	}

	testEvaluate(c, s.Tree, context, []string{
		"object01",
		"object03",
		"object07",
	})
}

func (s *DecisionSuite) Test_Context2(c *C) {
	context := TestContext{
		"geo_code":      "US",
		"platform":      "iOS",
		"device.gender": "male",
	}
	testEvaluate(c, s.Tree, context, []string{
		"object01",
		"object03",
	})
}

func (s *DecisionSuite) Test_Context3(c *C) {
	context := TestContext{
		"geo_code":      "US",
		"platform":      "iOS",
		"device.gender": "female",
	}
	testEvaluate(c, s.Tree, context, []string{
		"object01",
		"object03",
		"object07",
	})
}

func (s *DecisionSuite) Test_Context4(c *C) {
	context := TestContext{
		"geo_code": "US",
	}
	testEvaluate(c, s.Tree, context, []string{
		"object01",
		"object06",
	})
}

func (s *DecisionSuite) Test_Context5(c *C) {
	context := TestContext{
		"geo_code": "US",
		"platform": "iOS",
	}
	testEvaluate(c, s.Tree, context, []string{
		"object01",
		"object03",
		"object07",
	})
}

func (s *DecisionSuite) Test_Context6(c *C) {
	context := TestContext{
		"geo_code":      "US",
		"platform":      "iOS",
		"device.gender": "female",
	}
	testEvaluate(c, s.Tree, context, []string{
		"object01",
		"object03",
		"object07",
	})
}

func (s *DecisionSuite) Test_Context7(c *C) {
	context := TestContext{
		"geo_code":         "US",
		"platform":         "Android",
		"device.age_group": "60",
	}
	testEvaluate(c, s.Tree, context, []string{
		"object01",
		"object05",
		"object06",
	})
}

func (s *DecisionSuite) Test_Context8(c *C) {
	context := TestContext{
		"geo_code": "CA",
		"platform": "Android",
	}
	testEvaluate(c, s.Tree, context, []string{
		"object04",
	})
}

func (s *DecisionSuite) Test_Context9(c *C) {
	context := TestContext{
		"geo_code":      "CA",
		"platform":      "iOS",
		"device.gender": "female",
	}
	testEvaluate(c, s.Tree, context, []string{
		"object02",
	})
}

func (s *DecisionSuite) Test_Ages(c *C) {
	s.Objects = map[string][]string{
		"object01": []string{
			`geo_code matches '^(.*,)?(US)$'`,
		},
		"object02": []string{
			`platform = 'iOS'`,
		},
		"object03": []string{
			`platform = 'iOS'`,
			`geo_code matches '^(.*,)?(US)$'`,
		},
		"object04": []string{
			`platform = 'Android'`,
		},
		"object05": []string{
			`platform = 'Android'`,
			`device.age_group = "50"`,
			`geo_code matches '^(.*,)?(US)$'`,
		},
		"object06": []string{
			`device.age_group = "13-17"`,
			`geo_code matches '^(.*,)?(US)$'`,
		},
		"object07": []string{
			`geo_code matches '^(.*,)?(US)$'`,
			`platform = 'iOS'`,
		},
	}

	tree, err := NewTree(s.Objects)
	c.Assert(err, IsNil)

	f, err := os.Create("decision-ages.dot")
	c.Assert(err, IsNil)
	defer f.Close()
	tree.Graph(f)

	testEvaluate(
		c,
		tree,
		TestContext{
			"geo_code": "US",
		}, []string{
			"object01",
		})
	testEvaluate(
		c,
		tree,
		TestContext{
			"geo_code":         "US",
			"device.age_group": "13-17",
		}, []string{
			"object01",
			"object06",
		})
	testEvaluate(
		c,
		tree,
		TestContext{
			"geo_code":         "US",
			"device.age_group": "18-34",
		}, []string{
			"object01",
		})
	testEvaluate(
		c,
		tree,
		TestContext{
			"geo_code":         "US",
			"device.age_group": "35-49",
		}, []string{
			"object01",
		})
	testEvaluate(
		c,
		tree,
		TestContext{
			"geo_code":         "US",
			"device.age_group": "50",
		}, []string{
			"object01",
		})

}

func (s *DecisionSuite) Test_NotEquals(c *C) {
	objects := map[string][]string{
		"object01": []string{
			"app != 'A'",
		},
	}

	tree, err := NewTree(objects)
	c.Assert(err, IsNil)

	testEvaluate(
		c,
		tree,
		TestContext{
			"app": "A",
		}, []string{})

	testEvaluate(
		c,
		tree,
		TestContext{
			"app": "B",
		}, []string{
			"object01",
		})

	testEvaluate(
		c,
		tree,
		TestContext{
			"country": "US",
		}, []string{
			"object01",
		})
}

/*
func (s *DecisionSuite) Test_Find(c *C) {
	for object, expressions := range s.Objects {

		path, found := s.Tree.Find(object)
		c.Assert(found, Equals, true)

		pathMatch := []string{}
		notMatch := []string{}
		for _, pathStr := range path {
			matched := false
			for _, expStr := range expressions {
				exp, err := NewExpression(expStr)
				c.Assert(err, IsNil)
				if pathStr == exp.String() {
					pathMatch = append(pathMatch, pathStr)
					matched = true
				}
			}

			if !matched {
				notMatch = append(notMatch, pathStr)
			}
		}

		c.Assert(len(pathMatch), Equals, len(expressions))
		c.Assert(len(pathMatch)+len(notMatch), Equals, len(path))

		for _, expstr := range expressions {
			expression, err := NewExpression(expstr)
			c.Assert(err, IsNil)
			c.Assert(path, Contains, expression.String())
		}

		for _, notStr := range notMatch {
			check := strings.Contains(notStr, "!=") || strings.Contains(notStr, "NOT")
			c.Assert(check, Equals, true)
		}
	}
}
*/

func testEvaluate(c *C, tree *Tree, context exp.Context, expected []string) {
	logger := &TestLogger{make([]string, 0, 10)}

	list, err := tree.Evaluate(context, logger)
	c.Assert(err, IsNil)

	c.Assert(list, HasLen, len(expected))

	for _, object := range expected {
		c.Assert(list, Contains, object)
	}
}

func (s *DecisionSuite) Test_Matches(c *C) {
	exp, err := NewExpression(`test matches '^(hello|world)$'`)
	c.Assert(err, IsNil)

	{
		ctx := TestContext{}
		result, err := exp.Evaluate(ctx)
		c.Assert(err, IsNil)
		c.Assert(result, Equals, false)
	}
	{
		ctx := TestContext{
			"test": "qwerty",
		}
		result, err := exp.Evaluate(ctx)
		c.Assert(err, IsNil)
		c.Assert(result, Equals, false)
	}
	{
		ctx := TestContext{
			"test": "hello",
		}
		result, err := exp.Evaluate(ctx)
		c.Assert(err, IsNil)
		c.Assert(result, Equals, true)
	}
	{
		ctx := TestContext{
			"test": "world",
		}
		result, err := exp.Evaluate(ctx)
		c.Assert(err, IsNil)
		c.Assert(result, Equals, true)
	}
}

func (s *DecisionSuite) Test_NotMatches(c *C) {
	exp, err := NewExpression(`not test matches '^(hello|world)$'`)
	c.Assert(err, IsNil)

	{
		ctx := TestContext{}
		result, err := exp.Evaluate(ctx)
		c.Assert(err, IsNil)
		c.Assert(result, Equals, true)
	}
	{
		ctx := TestContext{
			"test": "qwerty",
		}
		result, err := exp.Evaluate(ctx)
		c.Assert(err, IsNil)
		c.Assert(result, Equals, true)
	}
	{
		ctx := TestContext{
			"test": "hello",
		}
		result, err := exp.Evaluate(ctx)
		c.Assert(err, IsNil)
		c.Assert(result, Equals, false)
	}
	{
		ctx := TestContext{
			"test": "world",
		}
		result, err := exp.Evaluate(ctx)
		c.Assert(err, IsNil)
		c.Assert(result, Equals, false)
	}
}

func (s *DecisionSuite) Test_NextExpression1(c *C) {
	expressions := []string{
		`geo_code matches '^(.*,)?(US)$'`,
		`platform = 'iOS'`,
		`device.age_group != "13-17"`,
	}

	{
		match, err := NewExpression(`platform = 'iOS'`)
		c.Assert(err, IsNil)

		expression, expressions2, err := nextExpression(expressions, match)
		c.Assert(err, IsNil)
		c.Assert(expression.String(), Equals, match.String())
		c.Assert(expressions2, HasLen, 2)
		c.Assert(expressions2, Contains, `geo_code matches '^(.*,)?(US)$'`)
		c.Assert(expressions2, Contains, `device.age_group != "13-17"`)
	}
	{
		match, err := NewExpression(`geo_code matches '^(.*,)?(US)$'`)
		c.Assert(err, IsNil)

		expression, expressions2, err := nextExpression(expressions, match)
		c.Assert(err, IsNil)
		c.Assert(expression.String(), Equals, match.String())
		c.Assert(expressions2, HasLen, 2)
		c.Assert(expressions2, Contains, `platform = 'iOS'`)
		c.Assert(expressions2, Contains, `device.age_group != "13-17"`)
	}
	{
		match, err := NewExpression(`device.age_group != "13-17"`)
		c.Assert(err, IsNil)

		expression, expressions2, err := nextExpression(expressions, match)
		c.Assert(err, IsNil)
		c.Assert(expression.String(), Equals, match.String())
		c.Assert(expressions2, HasLen, 2)
		c.Assert(expressions2, Contains, `platform = 'iOS'`)
		c.Assert(expressions2, Contains, `geo_code matches '^(.*,)?(US)$'`)
	}
	{
		first, err := NewExpression(`geo_code matches '^(.*,)?(US)$'`)
		c.Assert(err, IsNil)

		expression, expressions2, err := nextExpression(expressions, nil)
		c.Assert(err, IsNil)
		c.Assert(expression.String(), Equals, first.String())
		c.Assert(expressions2, HasLen, 2)
		c.Assert(expressions2, Contains, `platform = 'iOS'`)
		c.Assert(expressions2, Contains, `device.age_group != "13-17"`)
	}
	{
		expression, expressions2, err := nextExpression([]string{}, nil)
		c.Assert(err, IsNil)
		c.Assert(expression, IsNil)
		c.Assert(expressions2, HasLen, 0)
	}
}

func (s *DecisionSuite) Test_NextExpression2(c *C) {
	expressions := []string{
		`geo_code matches '^(.*,)?(US)$'`,
		`device.age_group != "13-17"`,
	}

	{
		match, err := NewExpression(`platform = 'iOS'`)
		c.Assert(err, IsNil)

		expression, expressions2, err := nextExpression(expressions, match)
		c.Assert(err, IsNil)
		c.Assert(expression, IsNil)
		c.Assert(expressions2, HasLen, 2)
		c.Assert(expressions2, Contains, `geo_code matches '^(.*,)?(US)$'`)
		c.Assert(expressions2, Contains, `device.age_group != "13-17"`)
	}
	{
		match, err := NewExpression(`geo_code matches '^(.*,)?(US)$'`)
		c.Assert(err, IsNil)

		expression, expressions2, err := nextExpression(expressions, match)
		c.Assert(err, IsNil)
		c.Assert(expression.String(), Equals, match.String())
		c.Assert(expressions2, HasLen, 1)
		c.Assert(expressions2, Contains, `device.age_group != "13-17"`)
	}
	{
		match, err := NewExpression(`device.age_group != "13-17"`)
		c.Assert(err, IsNil)

		expression, expressions2, err := nextExpression(expressions, match)
		c.Assert(err, IsNil)
		c.Assert(expression.String(), Equals, match.String())
		c.Assert(expressions2, HasLen, 1)
		c.Assert(expressions2, Contains, `geo_code matches '^(.*,)?(US)$'`)
	}
}

type containsChecker struct {
	*CheckerInfo
}

var Contains Checker = &containsChecker{
	&CheckerInfo{Name: "Contains", Params: []string{"obtained", "expected"}},
}

func (checker *containsChecker) Check(params []interface{}, names []string) (result bool, error string) {
	defer func() {
		if v := recover(); v != nil {
			result = false
			error = fmt.Sprint(v)
		}
	}()

	return stringInSlice(params[1].(string), params[0].([]string)), ""
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (s *DecisionSuite) TestLogicalOrEvaluation(c *C) {
	ctx := testContext{}

	// Test true or true = true
	expA, err := NewExpression("5 > 3")
	c.Assert(err, IsNil)
	expB, err := NewExpression("5 != 6")
	c.Assert(err, IsNil)

	exp := &exp.LogicalExpression{
		Left:    expA,
		Right:   expB,
		Logical: exp.Or,
	}

	result, err := exp.Evaluate(ctx)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, true)

	// Test false or true = true
	expA, err = NewExpression("5 < 3")
	c.Assert(err, IsNil)
	expB, err = NewExpression("5 != 6")
	c.Assert(err, IsNil)

	exp.Left = expA
	exp.Right = expB
	result, err = exp.Evaluate(ctx)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, true)

	// Test true or false = true
	expA, err = NewExpression("5 > 3")
	c.Assert(err, IsNil)
	expB, err = NewExpression("5 == 6")
	c.Assert(err, IsNil)

	exp.Left = expA
	exp.Right = expB
	result, err = exp.Evaluate(ctx)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, true)

	// Test false or false = false
	expA, err = NewExpression("5 < 3")
	c.Assert(err, IsNil)
	expB, err = NewExpression("5 == 6")
	c.Assert(err, IsNil)

	exp.Left = expA
	exp.Right = expB
	result, err = exp.Evaluate(ctx)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, false)
}

func (s *DecisionSuite) TestLogicalAndEvaluation(c *C) {
	ctx := testContext{}

	// Test true or true = true
	expA, err := NewExpression("5 > 3")
	c.Assert(err, IsNil)
	expB, err := NewExpression("5 != 6")
	c.Assert(err, IsNil)

	exp := &exp.LogicalExpression{
		Left:    expA,
		Right:   expB,
		Logical: exp.And,
	}

	result, err := exp.Evaluate(ctx)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, true)

	// Test false or true = true
	expA, err = NewExpression("5 < 3")
	c.Assert(err, IsNil)
	expB, err = NewExpression("5 != 6")
	c.Assert(err, IsNil)

	exp.Left = expA
	exp.Right = expB
	result, err = exp.Evaluate(ctx)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, false)

	// Test true or false = true
	expA, err = NewExpression("5 > 3")
	c.Assert(err, IsNil)
	expB, err = NewExpression("5 == 6")
	c.Assert(err, IsNil)

	exp.Left = expA
	exp.Right = expB
	result, err = exp.Evaluate(ctx)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, false)

	// Test false or false = false
	expA, err = NewExpression("5 < 3")
	c.Assert(err, IsNil)
	expB, err = NewExpression("5 == 6")
	c.Assert(err, IsNil)

	exp.Left = expA
	exp.Right = expB
	result, err = exp.Evaluate(ctx)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, false)
}

func (s *DecisionSuite) TestNegationEvaluation(c *C) {
	ctx := testContext{}

	expA, err := NewExpression("5 > 3")
	c.Assert(err, IsNil)
	c.Assert(expA, NotNil)

	exp := &exp.NegationExpression{
		Expression: expA,
	}
	result, err := exp.Evaluate(ctx)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, false)

	expA, err = NewExpression("5 < 3")
	c.Assert(err, IsNil)
	c.Assert(expA, NotNil)
	exp.Expression = expA
	result, err = exp.Evaluate(ctx)
	c.Assert(err, IsNil)
	c.Assert(result, Equals, true)
}

type testContext struct {
}

func (t testContext) Get(key string) (interface{}, bool) {
	return true, true
}
