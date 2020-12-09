package day8

import (
	"adventOfCode2020/computer"
	"adventOfCode2020/util"
	"fmt"
)

// Run day seven of advent
func Run(input <-chan string) {
	instructions := util.ChannelToSlice(input)
	//	fmt.Println("Day8 -- Part1: ", computer.New(instructions).ExecuteProgram())
	fmt.Println("Day8 -- Part2: ", computer.New(instructions).ExecuteSelfCorrectingProgram())
}
