package decide

import (
	"bytes"
	"fmt"
	"sort"
)

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
	sort.Sort(t)
}

func (t FrequencySorter) SortReverse(list []string) {
	t.list = list
	sort.Sort(sort.Reverse(t))
}

func (t FrequencySorter) FrequencyList() []string {
	list := make([]string, 0, len(t.frequencies))
	for key, _ := range t.frequencies {
		list = append(list, key)
	}
	t.SortReverse(list)
	return list
}

func (t FrequencySorter) AddToFrequencies(str string) {
	t.frequencies[str]++
}

func (t FrequencySorter) AddValue(str string, count int) {
	t.frequencies[str] = count
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
