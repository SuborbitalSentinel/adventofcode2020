package day15

import "fmt"

// Run executes day15 of advent
func Run() {
	part1Test := []int{0, 3, 6}
	fmt.Println("Day 15 -- part1 test:", calculateLastSpoken(part1Test, 2020))
	part1 := []int{1, 0, 15, 2, 10, 13}
	fmt.Println("Day 15 -- part 1:", calculateLastSpoken(part1, 2020))

	part2Test := []int{1, 3, 2}
	fmt.Println("Day 15 -- part2 test:", calculateLastSpoken(part2Test, 30000000))
	part2 := []int{1, 0, 15, 2, 10, 13}
	fmt.Println("Day 15 -- part 2:", calculateLastSpoken(part2, 30000000))
}

func calculateLastSpoken(input []int, rounds int) int {
	spokenNumbers := map[int][]int{}
	// Preamble to spoken game
	for i, n := range input {
		spokenNumbers[n] = []int{i}
	}
	spokenNumbers[0] = append(spokenNumbers[0], len(input))
	lastSpoken := 0
	for i := len(input) + 1; i < rounds; i++ {
		//fmt.Printf("Round: %d; Last Spoken: %d\n", i+1, lastSpoken)
		spokenMemory, ok := spokenNumbers[lastSpoken]
		if ok && len(spokenMemory) == 1 {
			spokenNumbers[0] = append(spokenNumbers[0], i)
			lastSpoken = 0
			continue
		}
		if ok && len(spokenMemory) > 1 {
			lastSpoken = spokenMemory[len(spokenMemory)-1] - spokenMemory[len(spokenMemory)-2]
			spokenNumbers[lastSpoken] = append(spokenNumbers[lastSpoken], i)
			continue
		}
	}
	return lastSpoken
}
