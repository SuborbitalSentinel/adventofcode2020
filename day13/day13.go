package day13

import (
	"fmt"
	"strconv"
)

// Run day thirteen of advent
func Run() {
	fmt.Println("Day13 -- Part1: ", part1())
	fmt.Println("Day13 -- Part2: ", part2())
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

type pair struct {
	Offset int
	ID     int
}

// Math trick problems are not fun to me...
// Something something Chinese Remainder Theorm
func part2() int {
	busIds := [...]string{"29", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "37", "x", "x", "x", "x", "x", "467", "x", "x", "x", "x", "x", "x", "x", "23", "x", "x", "x", "x", "13", "x", "x", "x", "17", "x", "19", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "443", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "x", "41"}
	buses := []pair{}
	for i, bus := range busIds {
		if bus != "x" {
			busID, _ := strconv.Atoi(bus)
			buses = append(buses, pair{Offset: i, ID: busID})
		}
	}
	time, step := 0, 1
	for _, p := range buses {
		for (time+p.Offset)%p.ID != 0 {
			time += step
		}
		step *= p.ID
	}
	return time
}
