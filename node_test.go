package decide

import (
	"fmt"
	exp "github.com/sjhitchner/go-decide/expression"
	. "gopkg.in/check.v1"
	"os"
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
	if err != nil {
		c.Fatal(err)
	}

	f, err := os.Create("decision.dot")
	if err != nil {
		c.Fatal(err)
	}
	defer f.Close()
	tree.Graph(f)

	s.Tree = tree
}

func (s *DecisionSuite) Test_Context1(c *C) {
	context := TestContext{
		"geo_code": "US",
	}

	logger := &TestLogger{make([]string, 0, 10)}

	list, err := s.Tree.Evaluate(context, logger)
	c.Assert(err, IsNil)

	c.Assert(list, HasLen, 1)
	c.Assert(list, Contains, "object01")
}

func (s *DecisionSuite) Test_Context2(c *C) {
	context := TestContext{
		"geo_code": "US",
		"platform": "iOS",
	}

	logger := &TestLogger{make([]string, 0, 10)}

	list, err := s.Tree.Evaluate(context, logger)
	c.Assert(err, IsNil)

	c.Assert(list, HasLen, 2)
	c.Assert(list, Contains, "object01")
	c.Assert(list, Contains, "object03")
}

func (s *DecisionSuite) Test_Context3(c *C) {
	context := TestContext{
		"geo_code":      "US",
		"platform":      "iOS",
		"device.gender": "female",
	}
	logger := &TestLogger{make([]string, 0, 10)}

	list, err := s.Tree.Evaluate(context, logger)
	c.Assert(err, IsNil)

	c.Assert(list, HasLen, 3)
	c.Assert(list, Contains, "object01")
	c.Assert(list, Contains, "object03")
	c.Assert(list, Contains, "object07")
}

func (s *DecisionSuite) Test_Context4(c *C) {
	context := TestContext{
		"geo_code":         "US",
		"platform":         "iOS",
		"device.age_group": "60",
	}
	logger := &TestLogger{make([]string, 0, 10)}

	list, err := s.Tree.Evaluate(context, logger)
	c.Assert(err, IsNil)

	c.Assert(list, HasLen, 2)
	c.Assert(list, Contains, "object01")
	c.Assert(list, Contains, "object03")
}

func (s *DecisionSuite) Test_Context5(c *C) {
	context := TestContext{
		"geo_code":         "US",
		"platform":         "Android",
		"device.age_group": "60",
	}
	logger := &TestLogger{make([]string, 0, 10)}

	list, err := s.Tree.Evaluate(context, logger)
	c.Assert(err, IsNil)

	c.Assert(list, HasLen, 3)
	c.Assert(list, Contains, "object01")
	c.Assert(list, Contains, "object05")
	c.Assert(list, Contains, "object06")
}

func (s *DecisionSuite) Test_Context6(c *C) {
	context := TestContext{
		"geo_code": "CA",
		"platform": "Android",
	}
	logger := &TestLogger{make([]string, 0, 10)}

	list, err := s.Tree.Evaluate(context, logger)
	c.Assert(err, IsNil)

	c.Assert(list, HasLen, 1)
	c.Assert(list, Contains, "object04")
}

func (s *DecisionSuite) Test_Context7(c *C) {
	context := TestContext{
		"geo_code":      "CA",
		"platform":      "iOS",
		"device.gender": "female",
	}
	logger := &TestLogger{make([]string, 0, 10)}

	list, err := s.Tree.Evaluate(context, logger)
	c.Assert(err, IsNil)

	c.Assert(list, HasLen, 1)
	c.Assert(list, Contains, "object02")
}

/*
(device.device_type matches '(?i)(Phone)')
(device.device_type matches '(?i)(Tablet)')
((geo_code matches '^(.*,)?(GB)$'))
(device.age_group != "13-17" and (geo_code matches '^(.*,)?(US)$'))
(app.genre matches '^(Productivity)$')

((geo_code matches '^(.*,)?(US)$'))
((platform = 'iOS'))
((platform = 'iOS') and (geo_code matches '^(.*,)?(US)$'))
((platform = 'Android'))
((platform = 'Android') and (geo_code matches '^(.*,)?(US)$'))
(device.age_group != "13-17" and (geo_code matches '^(.*,)?(US)$'))

*/

/*
((platform = 'Web Android' or platform = 'Android' or platform = 'Web Desktop') and (geo_code matches '^(.*,)?(FR)$'
))
((geo_code matches '^(.*,)?(CA)$'))
((platform = 'iOS' or platform = 'Web Desktop' or platform = 'Web iOS') and (geo_code matches '^(.*,)?(FR)$'))
((geo_code matches '^(.*,)?(HK)$'))
((platform = 'Web Android' or platform = 'Android' or platform = 'Web Desktop'))
((geo_code matches '^(.*,)?(02,AU|04,AU)$'))
((platform = 'Android') and (geo_code matches '^(.*,)?(MX)$'))
*/

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
