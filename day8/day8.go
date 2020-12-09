package day8

import (
	"adventOfCode2020/computer"
	"adventOfCode2020/util"
	"fmt"
)

// Run day seven of advent
func Run(input <-chan string) {
	computer := computer.New(util.ChannelToSlice(input))
	fmt.Println("Day8 -- Part1: ", computer.ExecuteProgram())
	//fmt.Println("Day8 -- Part2: ", 2)
}
