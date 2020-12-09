package types

import (
	"adventOfCode2020/util"
	"strconv"
)

// Register represents the contents of a memory location
type Register struct {
	Command  string
	Executed bool
	Mutated  bool
}

// Instruction returns the instruction of the command
func (r *Register) Instruction() string {
	c := util.RegexMap(`(?P<Instruction>\w+) (?P<Direction>(\+|-))(?P<Count>\d+)$`, r.Command)
	return c["Instruction"]
}

// Execute runs the command
func (r *Register) Execute() (pcOffset, accumulator int) {
	r.Executed = true
	c := util.RegexMap(`(?P<Instruction>\w+) (?P<Direction>(\+|-))(?P<Count>\d+)$`, r.Command)
	switch c["Instruction"] {
	case "nop":
		return 1, 0
	case "acc":
		inc, _ := strconv.Atoi(c["Direction"] + c["Count"])
		return 1, inc
	case "jmp":
		inc, _ := strconv.Atoi(c["Direction"] + c["Count"])
		return inc, 0
	}
	panic("Failed To Execute Register!")
}

// Revert undoes the Execute() function
func (r *Register) Revert() (pcOffset, accumulator int) {
	r.Executed = false
	c := util.RegexMap(`(?P<Instruction>\w+) (?P<Direction>(\+|-))(?P<Count>\d+)$`, r.Command)
	switch c["Instruction"] {
	case "nop":
		return -1, 0
	case "acc":
		inc, _ := strconv.Atoi(c["Direction"] + c["Count"])
		return -1, -1 * inc
	case "jmp":
		inc, _ := strconv.Atoi(c["Direction"] + c["Count"])
		return -1 * inc, 0
	}
	panic("Failed To Revert Register!")
}

// Mutate will flip nop and jmp instructions
func (r *Register) Mutate() {
	r.Mutated = true
	c := util.RegexMap(`(?P<Instruction>\w+) (?P<Direction>(\+|-))(?P<Count>\d+)$`, r.Command)
	switch c["Instruction"] {
	case "jmp":
		r.Command = "nop " + c["Direction"] + c["Count"]
	case "nop":
		r.Command = "jmp " + c["Direction"] + c["Count"]
	}
}
