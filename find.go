package decide

import (
	"fmt"
)

type stack struct {
	s []string
}

func (t *stack) Push(v string) {
	t.s = append(t.s, v)
}

func (t *stack) Pop() {
	l := len(t.s)
	if l != 0 {
		t.s = t.s[:l-1]
	}
}

func (t *stack) Len() int {
	return len(t.s)
}

func (t stack) String() string {
	return fmt.Sprintf("%v", t.s)
}

func (t Tree) Find(payload string) ([]string, bool) {
	path := &stack{make([]string, 0, 20)}
	found := t.root.Find(0, path, payload)
	return path.s, found
}

func (t Node) Find(depth int, path *stack, object string) bool {
	path.Push(t.Expression.String())

	if t.Contains(object) {
		return true
	}

	if len(t.True) > 0 {
		for _, trueNode := range t.True {
			if trueNode.Find(depth+1, path, object) {
				return true
			}
		}
	}

	path.Pop()

	if t.False != nil && t.False.Find(depth+1, path, object) {
		return true
	}

	return false
}

func (t Node) Contains(object string) bool {
	for _, p := range t.Payload {
		if p == object {
			return true
		}
	}
	return false
}
