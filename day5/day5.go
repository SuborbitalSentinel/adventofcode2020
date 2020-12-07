package day5

import (
	"adventOfCode2020/util"
	"fmt"
	"sort"
)

// Run day five of advent
func Run(input <-chan string) {
	var planeRows [128]int
	for i := range planeRows {
		planeRows[i] = i
	}
	planeCols := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	boardingPasses := util.ChannelToSlice(input)

	fmt.Println("Day5 -- Part1: ", maxSeatID(planeRows, planeCols, boardingPasses))
	fmt.Println("Day5 -- Part2: ", findMissingSeatID(planeRows, planeCols, boardingPasses))
}

func findMissingSeatID(planeRows [128]int, planeCols [8]int, boardingPasses []string) int {
	var seatIDs []int
	for _, boardingPass := range boardingPasses {
		seatIDs = append(seatIDs, calculateSeatID(findPlaneSeat(boardingPass, planeRows[:], planeCols[:])))
	}
	sort.Ints(seatIDs)
	for i, seatID := range seatIDs {
		if i == 0 || i == len(seatIDs)-1 {
			continue
		}
		if possibleID := seatID - 1; possibleID != seatIDs[i-1] {
			return possibleID
		}
		if possibleID := seatID + 1; possibleID != seatIDs[i+1] {
			return possibleID
		}
	}
	return -1
}

func maxSeatID(planeRows [128]int, planeCols [8]int, boardingPasses []string) int {
	maxID := 0
	for _, boardingPass := range boardingPasses {
		if ID := calculateSeatID(findPlaneSeat(boardingPass, planeRows[:], planeCols[:])); ID > maxID {
			maxID = ID
		}
	}
	return maxID
}

func minSeatID(planeRows [128]int, planeCols [8]int, boardingPasses []string) int {
	minID := (127 * 8) + 7
	for _, boardingPass := range boardingPasses {
		if ID := calculateSeatID(findPlaneSeat(boardingPass, planeRows[:], planeCols[:])); ID < minID {
			minID = ID
		}
	}
	return minID
}

func calculateSeatID(row, col int) int {
	return (row * 8) + col
}

func findPlaneSeat(input string, planeRows, planeCols []int) (row, col int) {
	row = binarySearch(input[:7], 'F', 'B', planeRows)
	col = binarySearch(input[7:], 'L', 'R', planeCols)
	return
}

func binarySearch(decision string, lower, upper byte, target []int) int {
	if len(decision) == 0 {
		return target[0]
	}
	mid := len(target) / 2
	if decision[0] == lower {
		return binarySearch(decision[1:], lower, upper, target[:mid])
	}
	if decision[0] == upper {
		return binarySearch(decision[1:], lower, upper, target[mid:])
	}
	return -1
}
