package day11

import (
	"adventOfCode2020/types"
	"adventOfCode2020/util"
	"fmt"
	"strings"
)

func padInput(input []string) []string {
	for i, row := range input {
		input[i] = "." + row + "."
	}
	rowLength := len(input[0])
	input = append([]string{strings.Repeat(".", rowLength)}, input...)
	input = append(input, strings.Repeat(".", rowLength))
	return input
}

// Run day eleven of advent
func Run(input <-chan string) {
	planeSeats := util.ChannelToSlice(input)
	planeSeats = padInput(planeSeats)

	//fmt.Println("Day11 -- Part1: ", part1(planeSeats))
	fmt.Println("Day11 -- Part2: ", part2(planeSeats))
}

func part1(input []string) int {
	plane := types.NewPlane(input)
	for plane.NextSeatingArrangement() > 0 {
		plane.Print()
	}
	return plane.CountOccupiedSeats()
}

func part2(input []string) int {
	plane := types.NewPlane(input)
	plane.Print()
	for plane.NextSeatingArrangement() > 0 {
		plane.Print()
	}
	return plane.CountOccupiedSeats()
}
