package computer

import (
	"adventOfCode2020/types"
)

// Computer represents a type that can run programs
type Computer struct {
	Accumulator            int
	ProgramCounter         int
	Memory                 []types.Register
	CallStack              types.RegisterStack
	HasAttemptedCorrection bool
}

// New returns a new computer type
func New(instructionSet []string) *Computer {
	var registers []types.Register
	for _, instruction := range instructionSet {
		register := types.Register{Command: instruction, Executed: false, CurrentlyMutated: false, HasEverBeenMutated: false}
		registers = append(registers, register)
	}
	return &Computer{
		Accumulator:            0,
		ProgramCounter:         0,
		Memory:                 registers,
		HasAttemptedCorrection: false,
	}
}

// CleanMemory scans the computer memory and calls Mutate on any memory location that has been mutated
func (c *Computer) CleanMemory() {
	c.HasAttemptedCorrection = false
	for _, reg := range c.Memory {
		if reg.CurrentlyMutated {
			reg.Mutate()
		}
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

// AttemptCorrection walks back the callstack until it finds a jmp, mutates it, then exits
func (c *Computer) AttemptCorrection() {
	c.HasAttemptedCorrection = true
	for true {
		c.RevertLastInstruction()
		if !c.CurrentInstruction().HasEverBeenMutated && (c.CurrentInstruction().Instruction() == "jmp" || c.CurrentInstruction().Instruction() == "nop") {
			c.CurrentInstruction().Mutate()
			break
		}
	}
}

// ExecuteSelfCorrectingProgram runs the program stored in the computer but tries to fix the infinite loop
func (c *Computer) ExecuteSelfCorrectingProgram() int {
	for true {
		if c.ProgramCounter > len(c.Memory)-1 {
			break
		}
		if c.CurrentInstruction().Executed {
			if c.HasAttemptedCorrection {
				c.CleanMemory()
			}
			c.AttemptCorrection()
		}
		c.ExecuteCurrentInstruction()
	}
	return c.Accumulator
}
