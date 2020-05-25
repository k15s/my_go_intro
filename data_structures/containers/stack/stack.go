package stack

// It’s also worth noting that Go provides container/list and container/ring
// packages for doubly-linked and circular lists, respectively. These classes
// operate on empty interfaces, just like this Stack type

// This stack is backed by an underlying array. Because it does not resize
// upon removals, we also track the index of the top of the stack in the array.
type Stack struct {
	items     []interface{}
	lastIndex int
}

func NewStack() *Stack {
	return &Stack{items: make([]interface{}, 0), lastIndex: -1}
}

// push `v` onto the stack
func (s *Stack) Push(v interface{}) {
	if s.lastIndex != -1 && s.lastIndex < len(s.items)-1 {
		s.items[s.lastIndex+1] = v
	} else {
		s.items = append(s.items, v)
	}
	s.lastIndex++
}

// Pop the top element off the stack
func (s *Stack) Pop() interface{} {
	e := s.items[s.lastIndex]
	s.lastIndex--
	return e
}

func (s *Stack) Peek() interface{} {
	return s.items[s.lastIndex]
}

func (s *Stack) Size() int {
	return s.lastIndex + 1
}

// Int-specific, type-safe wrapper
type IntStack struct {
	*Stack
}

func NewIntStack() *IntStack {
	return &IntStack{NewStack()}
}

func (s *IntStack) Push(i int) {
	s.Stack.Push(i)
}

func (s *IntStack) Pop() int {
	// Pop the top element from the stack and type assert it to int
	return s.Stack.Pop().(int)
}
