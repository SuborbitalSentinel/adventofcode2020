package computer

import (
	"adventOfCode2020/types"
)

// Computer represents a type that can run programs
type Computer struct {
	Accumulator    int
	ProgramCounter int
	Memory         []types.Register
	CallStack      types.RegisterStack
}

// New returns a new computer type
func New(instructionSet []string) *Computer {
	var registers []types.Register
	for _, instruction := range instructionSet {
		register := types.Register{Command: instruction, Executed: false}
		registers = append(registers, register)
	}
	return &Computer{
		Accumulator:    0,
		ProgramCounter: 0,
		Memory:         registers,
	}
}

// CurrentInstruction returns the register of the current ProgramCounter
func (c *Computer) CurrentInstruction() *types.Register {
	return &c.Memory[c.ProgramCounter]
}

// ExecuteCurrentInstruction evaluates the current instruction of the program counter
func (c *Computer) ExecuteCurrentInstruction() {
	c.CallStack.Push(c.CurrentInstruction())
	pc, acc := c.CurrentInstruction().Execute()
	c.ProgramCounter += pc
	c.Accumulator += acc
}

// RevertLastInstruction reverts the top instruction on the stack
func (c *Computer) RevertLastInstruction() {
	reg := c.CallStack.Pop()
	pc, acc := reg.Revert()
	c.ProgramCounter += pc
	c.Accumulator += acc
}

// ExecuteProgram runs the program stored in the computer
func (c *Computer) ExecuteProgram() int {
	for true {
		if c.CurrentInstruction().Executed {
			break
		}
		c.ExecuteCurrentInstruction()
	}
	return c.Accumulator
}
