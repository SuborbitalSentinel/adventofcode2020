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

	part1, part2 := parts(directions)
	fmt.Println("Day12 -- Part1: ", part1)
	fmt.Println("Day12 -- Part2: ", part2)
}

func parts(input []string) (part1, part2 int) {
	ns1 := types.NewNavigationSystem()
	ns2 := types.NewWaypointNavigationSystem()
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
		ns1.Move[direction](magnitude)
		ns2.Move[direction](magnitude)
	}
	return ns1.ManhattenDistance(), ns2.ManhattenDistance()
}
