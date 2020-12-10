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
		register := types.Register{Command: instruction, Executed: false, Mutated: false, TriedCorrection: false}
		registers = append(registers, register)
	}
	return &Computer{
		Accumulator:    0,
		ProgramCounter: 0,
		Memory:         registers,
	}
}

// Reset resets everything on the computer except for AttemptedCorrection
func (c *Computer) Reset() {
	c.ProgramCounter = 0
	c.Accumulator = 0
	c.CallStack.Clear()
	for i := range c.Memory {
		c.Memory[i].Executed = false
	}
}

// RevertAllMutations loops through the computers memory reverting any mutation that has been done
func (c *Computer) RevertAllMutations() {
	for i := range c.Memory {
		if c.Memory[i].Mutated {
			c.Memory[i].Mutate()
		}
	}
}

// IsProgramCounterInRange returns true if the program counter is no longer usable
func (c *Computer) IsProgramCounterInRange() bool {
	return 0 <= c.ProgramCounter && c.ProgramCounter <= len(c.Memory)-1
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

// FindPossiblyCorruptRegister walks back the callstack until it finds a jmp, mutates it, then exits
func (c *Computer) FindPossiblyCorruptRegister() *types.Register {
	for c.CallStack.Count() != 0 {
		if c.CurrentInstruction().IsPossiblyCorrupt() {
			return c.CurrentInstruction()
		}
		c.RevertLastInstruction()
	}
	panic("We popped the whole stack and didn't find a command to mutate")
}

// ExecuteSelfCorrectingProgram runs the program stored in the computer but tries to fix the infinite loop
func (c *Computer) ExecuteSelfCorrectingProgram() int {
	for c.IsProgramCounterInRange() {
		if c.CurrentInstruction().Executed {
			reg := c.FindPossiblyCorruptRegister()
			c.RevertAllMutations()
			c.Reset()
			reg.Mutate()
		}
		c.ExecuteCurrentInstruction()
	}
	return c.Accumulator
}
