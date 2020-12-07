package day1

import (
	"adventOfCode2020/util"
	"fmt"
)

// Run day one of advent
func Run(input <-chan string) {
	expenseReportValues := util.StringChannelToIntSlice(input)

	fmt.Println("Day1 -- Part1: ", part1(expenseReportValues))
	fmt.Println("Day1 -- Part2: ", part2(expenseReportValues))
}

func part1(input []int) int {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i != j && input[i]+input[j] == 2020 {
				return input[i] * input[j]
			}
		}
	}
	return -1
}

func part2(input []int) int {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			for k := 0; k < len(input); k++ {
				if i != j && i != k && input[i]+input[j]+input[k] == 2020 {
					return input[i] * input[j] * input[k]
				}
			}
		}
	}

	return -1
}
