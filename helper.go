package decisions

import (
	"bytes"
	"fmt"
	"sort"
)

type ExpressionPair struct {
	Expression string
	Count      int
}

type ExpressionPairList []ExpressionPair

func (t ExpressionPair) String() string {
	return fmt.Sprintf("%d: %s", t.Count, t.Expression)
}

func (t ExpressionPairList) Len() int {
	return len(t)
}

func (t ExpressionPairList) Less(i, j int) bool {
	return t[i].Count < t[j].Count
}

func (t ExpressionPairList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func NewExpressionPairList(objects map[Object][]string) ExpressionPairList {
	expression := make(map[string]int)

	for _, expressions := range objects {
		for _, exp := range expressions {
			expression[exp] += 1
		}
	}

	return rankExpressionByFrequency(expression)
}

func rankExpressionByFrequency(expressionFrequencies map[string]int) ExpressionPairList {
	el := make(ExpressionPairList, len(expressionFrequencies))

	i := 0
	for k, v := range expressionFrequencies {
		el[i] = ExpressionPair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(el))
	return el
}

// FrequencySorter
// Sorts Frequency Lists not thread-safe
type FrequencySorter struct {
	frequencies map[string]int
	list        []string
}

func NewFrequencySorter() FrequencySorter {
	return FrequencySorter{
		frequencies: make(map[string]int),
		list:        nil,
	}
}

func (t FrequencySorter) Sort(list []string) {
	t.list = list
	sort.Sort(sort.Reverse(t))
}

func (t FrequencySorter) FrequencyList() []string {
	list := make([]string, 0, len(t.frequencies))
	for key, _ := range t.frequencies {
		list = append(list, key)
	}
	t.Sort(list)
	return list
}

func (t FrequencySorter) AddToFrequencies(str string) {
	t.frequencies[str]++
}

func (t FrequencySorter) Len() int {
	return len(t.list)
}

func (t FrequencySorter) Less(i, j int) bool {
	if t.frequencies[t.list[i]] == t.frequencies[t.list[j]] {
		return t.list[i] < t.list[j]
	}
	return t.frequencies[t.list[i]] < t.frequencies[t.list[j]]
}

func (t FrequencySorter) Swap(i, j int) {
	t.list[i], t.list[j] = t.list[j], t.list[i]
}

func (t FrequencySorter) String() string {
	buf := &bytes.Buffer{}
	fmt.Fprintln(buf, "Frequency Table:")
	for key, val := range t.frequencies {
		fmt.Fprintf(buf, "% 4d: %s\n", val, key)
	}
	return buf.String()
}
