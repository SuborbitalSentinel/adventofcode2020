package types

// RegisterStack is a stack of registers
type RegisterStack struct {
	stack []*Register
}

// Dump returns all the current comamnd strings in the stack
func (s *RegisterStack) Dump() (out []string) {
	for _, reg := range s.stack {
		out = append(out, reg.Command)
	}
	return
}

// Peek returns a pointer to the top item in the stack, but does not remove it
func (s *RegisterStack) Peek() *Register {
	return s.stack[0]
}

// Push adds a register to the top of the stack
func (s *RegisterStack) Push(value *Register) {
	s.stack = append([]*Register{value}, s.stack...)
}

// Pop returns the top element and removes it from the stack
func (s *RegisterStack) Pop() *Register {
	r := s.stack[0]
	s.stack = s.stack[1:]
	return r
}

// Empty returns true if the stack is empty
func (s *RegisterStack) Empty() bool {
	return len(s.stack) == 0
}
