package day13

import "fmt"

// Run day thirteen of advent
func Run() {
	fmt.Println("Day13 -- Part1: ", part1())
	//fmt.Println("Day13 -- Part2: ", part2)
}

// Hard coded everything for part 1 because it seemed silly to parse it
func part1() int {
	earliestDeparture := 1000508
	busIds := [...]int{13, 17, 19, 23, 29, 37, 41, 443, 467}
	for i := earliestDeparture; i < 10000000; i++ {
		for _, bus := range busIds {
			if i%bus == 0 {
				return bus * (i - earliestDeparture)
			}
		}
	}
	panic("Bus not found")
}
