package decide

import (
	. "gopkg.in/check.v1"
	"math/rand"
)

func (s DecisionSuite) Test_FrequencySorterExpressions(c *C) {
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
			sorter.SortReverse(temp)
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
			sorter.SortReverse(temp)
			c.Assert(temp, DeepEquals, expected)
		}
	}
}

func (s *DecisionSuite) Test_FrequencySorterObjects(c *C) {

	objects := map[string]int{
		"object01": 3,
		"object02": 2,
		"object03": 1,
		"object04": 5,
		"object05": 3,
		"object06": 2,
		"object07": 4,
	}

	objectSorter := NewFrequencySorter()
	for object, count := range objects {
		objectSorter.AddValue(object, count)
	}

	sortedObjects := objectSorter.FrequencyList()
	c.Assert(sortedObjects, HasLen, 7)
	c.Assert(sortedObjects[0], Equals, "object04")
	c.Assert(sortedObjects[1], Equals, "object07")
	c.Assert(sortedObjects[2], Equals, "object05")
	c.Assert(sortedObjects[3], Equals, "object01")
	c.Assert(sortedObjects[4], Equals, "object06")
	c.Assert(sortedObjects[5], Equals, "object02")
	c.Assert(sortedObjects[6], Equals, "object03")
}
