package day3

import (
	"adventOfCode2020/util"
	"fmt"
)

// Run day two of advent
func Run(input <-chan string) {
	maze := util.ChannelToSlice(input)

	fmt.Println("Day3 -- Part1: ", traverseSlope(maze, 3, 1))
	fmt.Println("Day3 -- Part2: ", part2(maze))
}

func part2(maze []string) int {
	return traverseSlope(maze, 1, 1) * traverseSlope(maze, 3, 1) * traverseSlope(maze, 5, 1) * traverseSlope(maze, 7, 1) * traverseSlope(maze, 1, 2)
}

func traverseSlope(maze []string, run, rise int) (numberOfTrees int) {
	rowLen := len(maze[0])
	xPos := 0
	yPos := 0
	for yPos < len(maze)-rise {
		if xPos+run < rowLen {
			xPos += run
		} else {
			xPos = (xPos + run) - rowLen
		}
		yPos += rise

		if maze[yPos][xPos] == '#' {
			numberOfTrees++
		}

	}

	return
}
