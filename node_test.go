package decide

import (
	"fmt"
	exp "github.com/sjhitchner/go-decide/expression"
	. "gopkg.in/check.v1"
	"math/rand"
	"os"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type DecisionSuite struct{}

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

func (s *DecisionSuite) Test(c *C) {
	objects := map[string][]string{
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
	}

	tree, err := NewTree(objects)
	if err != nil {
		c.Fatal(err)
	}

	f, err := os.Create("decision.dot")
	if err != nil {
		c.Fatal(err)
	}
	defer f.Close()
	tree.Graph(f)

	{
		context := TestContext{
			"geo_code":         "US",
			"platform":         "iOS",
			"device.age_group": "60",
		}
		testEvaluate(c, tree, context, []string{
			"object01",
			"object03",
			"object06",
		})
	}

}

func (s *DecisionSuite) Test2(c *C) {
	objects := map[string][]string{
		"object01": []string{
			`is_test = true`,
			`geo_code matches '^(.*,)?(US)$'`,
		},
		"object02": []string{
			`is_test = true`,
			`platform = 'iOS'`,
		},
		"object03": []string{
			`is_test = false`,
			`platform = 'iOS'`,
			`geo_code matches '^(.*,)?(US)$'`,
		},
		"object04": []string{
			`is_test = false`,
			`platform = 'Android'`,
		},
		"object05": []string{
			`is_test = false`,
			`platform = 'Android'`,
			`geo_code matches '^(.*,)?(US)$'`,
		},
		"object06": []string{
			`is_test = false`,
			`device.age_group != "13-17"`,
			`geo_code matches '^(.*,)?(US)$'`,
		},
	}

	tree, err := NewTree(objects)
	if err != nil {
		c.Fatal(err)
	}

	f, err := os.Create("decision2.dot")
	if err != nil {
		c.Fatal(err)
	}
	defer f.Close()
	tree.Graph(f)

	{
		context := TestContext{
			"geo_code":         "US",
			"platform":         "iOS",
			"device.age_group": "60",
		}
		testEvaluate(c, tree, context, []string{
			"object01",
			"object03",
			"object06",
		})
	}
}

func testEvaluate(c *C, tree *Tree, context exp.Context, expected []string) {
	logger := &TestLogger{
		make([]string, 0, 10),
	}

	log, err := tree.Evaluate(context, logger)
	if err != nil {
		c.Fatal(err)
	}

	c.Assert(log, HasLen, len(expected))

	for _, obtained := range log {
		c.Assert(stringInSlice(obtained, expected), Equals, true)
	}

	c.Log(log)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
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

func (s DecisionSuite) TestFrequencySorter(c *C) {
	toSort := []string{
		`geo_code matches '^(.*,)?(US)$'`,
		`platform = 'iOS'`,
		`platform = 'iOS'`,
		`geo_code matches '^(.*,)?(US)$'`,
		`platform = 'Android'`,
		`platform = 'Android'`,
		`geo_code matches '^(.*,)?(US)$'`,
		`device.age_group != "13-17"`,
		`geo_code matches '^(.*,)?(US)$'`,
		`device.device_type matches '(?i)(Phone)'`,
		`device.device_type matches '(?i)(Phone)'`,
		`device.device_type matches '(?i)(Tablet)'`,
		`geo_code matches '^(.*,)?(GB)$'`,
		`device.age_group != "13-17"`,
		`geo_code matches '^(.*,)?(US)$'`,
		`app.genre matches '^(Productivity)$'`,
		`geo_code matches '^(.*,)?(US)$'`,
		`platform = 'iOS'`,
		`platform = 'iOS'`,
		`geo_code matches '^(.*,)?(US)$'`,
		`platform = 'Android'`,
		`platform = 'Android'`,
		`geo_code matches '^(.*,)?(US)$'`,
		`device.age_group != "13-17"`,
		`geo_code matches '^(.*,)?(US)$'`,
	}

	sorter := NewFrequencySorter()
	for _, str := range toSort {
		sorter.AddToFrequencies(str)
	}

	{
		// Test many sorted permatations
		expected := []string{
			`geo_code matches '^(.*,)?(US)$'`,
			`platform = 'iOS'`,
			`platform = 'Android'`,
			`device.age_group != "13-17"`,
			`device.device_type matches '(?i)(Phone)'`,
			`geo_code matches '^(.*,)?(GB)$'`,
			`device.device_type matches '(?i)(Tablet)'`,
			`app.genre matches '^(Productivity)$'`,
		}
		for i := 0; i < 100; i++ {
			var temp = make([]string, len(expected))
			copy(temp, expected)

			//Randomize List
			for i := range temp {
				j := rand.Intn(i + 1)
				temp[i], temp[j] = temp[j], temp[i]
			}
			sorter.Sort(temp)
			c.Assert(temp, DeepEquals, expected)
		}
	}
	{
		// Test many sorted permatations
		expected := []string{
			`geo_code matches '^(.*,)?(US)$'`,
			`platform = 'Android'`,
			`device.device_type matches '(?i)(Phone)'`,
			`app.genre matches '^(Productivity)$'`,
		}
		for i := 0; i < 100; i++ {
			var temp = make([]string, len(expected))
			copy(temp, expected)

			//Randomize List
			for i := range temp {
				j := rand.Intn(i + 1)
				temp[i], temp[j] = temp[j], temp[i]
			}
			sorter.Sort(temp)
			c.Assert(temp, DeepEquals, expected)
		}
	}
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
