package day10

import (
	"adventOfCode2020/util"
	"fmt"
	"sort"
)

// Run day ten of advent
func Run(input <-chan string) {
	jolts := util.StringChannelToIntSlice(input)

	fmt.Println("Day10 -- Part1: ", part1(jolts))
	fmt.Println("Day10 -- Part2: ", part2(jolts))
}

func part1(input []int) int {
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)

	singleJoltDifference := 0
	trippleJoltDifference := 0
	lastJoltRating := 0
	for _, jolt := range input {
		difference := jolt - lastJoltRating
		if difference == 1 {
			singleJoltDifference++
		}
		if difference == 3 {
			trippleJoltDifference++
		}
		lastJoltRating = jolt
	}
	return singleJoltDifference * trippleJoltDifference
}

// I am shamelessly looking up a god damned answer for this because I am not a mathematician
var tribofuckinacciSequence = [...]int{1, 1, 2, 4, 7, 13, 24, 44, 81, 149}

func getTribofuckinacci(input int) int {
	if input > len(tribofuckinacciSequence) {
		panic("Fuck you, input too big")
	}
	return tribofuckinacciSequence[input-1]
}

func part2(input []int) int {
	sort.Ints(input)
	maxJoltage := input[len(input)-1]
	input = append(input, 0, maxJoltage+3)
	sort.Ints(input)

	multiplier := 1
	currentRun := 1
	for _, jolt := range input {
		if contains(input, jolt+1) {
			currentRun++
		} else {
			multiplier *= getTribofuckinacci(currentRun)
			currentRun = 1
		}
	}
	return multiplier
}

func contains(input []int, value int) bool {
	for _, i := range input {
		if i == value {
			return true
		}
	}
	return false
}
