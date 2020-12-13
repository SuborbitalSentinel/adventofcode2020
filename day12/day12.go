package day12

import (
	"adventOfCode2020/types"
	"adventOfCode2020/util"
	"fmt"
	"strconv"
)

// Run day twelve of advent
func Run(input <-chan string) {
	directions := util.ChannelToSlice(input)

	fmt.Println("Day12 -- Part1: ", part1(directions))
	//fmt.Println("Day12 -- Part2: ", part2(planeSeats))
}

func part1(input []string) int {
	ns := types.NewNavigationSystem()
	d := map[rune]types.Direction{
		'N': types.North,
		'S': types.South,
		'E': types.East,
		'W': types.West,
		'L': types.Left,
		'R': types.Right,
		'F': types.Forward,
	}

	for _, command := range input {
		direction := d[rune(command[0])]
		magnitude, _ := strconv.Atoi(command[1:])
		ns.Move[direction](magnitude)
	}
	return ns.ManhattenDistance()
}
