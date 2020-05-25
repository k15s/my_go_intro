package set_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"my_go_intro/data_structures/containers/set"
)

// `go test` in this directory to run the tests

func Run() {
	s := set.Set()
	s.Add(1)
}

func genSet(e []int) set.SetWrapper {
	s := set.Set()
	for _, i := range e {
		s.Add(i)
	}
	return s
}

// go test set_test.go
func TestSet(t *testing.T) {
	s := set.Set()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	assert.Equal(t, len(s), 3, "Set has size 3")
	assert.True(t, s.Contains(1), "Set contains an element")
	assert.True(t, s.Contains(2), "Set contains an element")
	assert.False(t, s.Contains(-1), "Set does not contain an element")

	s.Add(1)
	s.Add(2)
	assert.Equal(t, len(s), 3, "Set has size 3")

	s.Discard(1)
	assert.Equal(t, len(s), 2, "Set has size 2")

	s.Discard(2)
	assert.Equal(t, len(s), 1, "Set has size 1")

	s.Clean()
	assert.Equal(t, len(s), 0, "Set is empty")
	assert.False(t, s.Contains(1), "Set contains an element")
	assert.False(t, s.Contains(2), "Set contains an element")
}

func TestStringSet(t *testing.T) {
	s := set.Set()
	s.Add("a")
	s.Add("b")
	s.Add("c")
	assert.Equal(t, len(s), 3, "Set has size 3")
	assert.True(t, s.Contains("a"), "Set contains an element")
	assert.True(t, s.Contains("b"), "Set contains an element")
	assert.False(t, s.Contains("asfasdf"), "Set does not contain an element")

	s.Add("a")
	s.Add("b")
	assert.Equal(t, len(s), 3, "Set has size 3")

	s.Discard("a")
	assert.Equal(t, len(s), 2, "Set has size 2")

	s.Discard("b")
	assert.Equal(t, len(s), 1, "Set has size 1")
	assert.True(t, s.Contains("c"), "Set contains an element")

	s.Clean()
	assert.Equal(t, len(s), 0, "Set is empty")
	assert.False(t, s.Contains("a"), "Set contains an element")
	assert.False(t, s.Contains("b"), "Set contains an element")
}

func TestContains(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5}
	s := genSet(elements)

	for _, i := range elements {
		if !s.Contains(i) {
			t.Errorf("The set should contain %d", i)
		}
	}

	if s.Contains(0) {
		t.Error("The set should not contain 0")
	}
}

func TestEquals(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5}
	elements2 := []int{6, 7, 8, 9, 10}
	elements3 := []int{6, 7, 8, 9, 10, 11, 12, 13}
	elements4 := []int{5, 4, 3, 2, 1}

	s := genSet(elements)
	a := genSet(elements)
	b := genSet(elements2)
	c := genSet(elements3)
	d := genSet(elements4)

	assert.True(t, s.Equals(a), "The sets should be equal")
	assert.True(t, s.Equals(d), "The sets should be equal")
	assert.False(t, s.Equals(b), "The sets should not be equal")
	assert.False(t, b.Equals(c), "The sets should not be equal")
}

func TestSubSet(t *testing.T) {
	elements := []int{1, 2, 3}
	elements2 := []int{1, 2, 3, 4, 5, 6, 7}
	elements3 := []int{4, 5, 6, 7, 8, 9, 10}

	a := genSet(elements)
	b := genSet(elements2)
	c := genSet(elements3)

	assert.True(t, a.Subset(b), "Should be a subset")
	assert.True(t, a.Subset(a), "Should be a subset")
	assert.True(t, set.Set().Subset(a), "Empty set is automatic subset")

	assert.False(t, a.Subset(c), "Should not be a subset")
	assert.False(t, b.Subset(c), "Should not be a subset")
}

func TestUnion(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5}
	elements2 := []int{6, 7, 8, 9, 10}

	a := genSet(elements)
	b := genSet(elements2)

	s := a.Union(b)
	assert.Equal(t, s, genSet(append(elements, elements2...)))
}

func TestIntersection(t *testing.T) {
	elements := []int{4, 5, 6, 7}
	elements2 := []int{1, 2, 3, 4, 5, 6, 7}
	elements3 := []int{4, 5, 6, 7, 8, 9, 10}

	a := genSet(elements)
	b := genSet(elements2)
	c := genSet(elements3)

	s := b.Intersection(c)

	assert.Equal(t, s, a, "")
}
