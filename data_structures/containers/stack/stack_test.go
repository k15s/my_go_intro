package stack_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"my_go_intro/data_structures/containers/stack"
)

func Run() {
	s := stack.NewStack()
	s.Push(1)
}

func genStack(e []int) *stack.Stack {
	s := stack.NewStack()
	for _, i := range e {
		(*s).Push(i)
	}
	return s
}

// go test stack_test.go
func TestIntStack(t *testing.T) {
	s := stack.NewIntStack()
	assert.Equal(t, (*s).Size(), 0, "Stack is empty")

	(*s).Push(1)
	(*s).Push(2)
	(*s).Push(3)
	assert.Equal(t, (*s).Size(), 3, "Stack has 3 elements")

	assert.Equal(t, (*s).Pop(), 3, "Stack pops off top element")
	assert.Equal(t, (*s).Pop(), 2, "Stack pops off top element")
	assert.Equal(t, (*s).Size(), 1, "Stack has 1 element")
	assert.Equal(t, (*s).Peek(), 1, "Stack peeks at top element")

	(*s).Push(2)
	(*s).Push(3)
	assert.Equal(t, (*s).Size(), 3, "Stack has 3 elements")
	assert.Equal(t, (*s).Peek(), 3, "Stack peeks at top element")

	assert.Equal(t, (*s).Pop(), 3, "Stack pops off top element")
	assert.Equal(t, (*s).Pop(), 2, "Stack pops off top element")
	assert.Equal(t, (*s).Pop(), 1, "Stack pops off top element")
	assert.Equal(t, (*s).Size(), 0, "Stack is empty")
}

func TestSize(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5}
	s := genStack(elements)
	assert.Equal(t, (*s).Size(), 5, "Stack size")
}

func TestPush(t *testing.T) {
	s := stack.NewStack()
	assert.Equal(t, (*s).Size(), 0, "Stack size")
	(*s).Push(1)
	(*s).Push(2)
	(*s).Push(3)
	assert.Equal(t, (*s).Size(), 3, "Stack size")
}
