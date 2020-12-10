package day9

import (
	"adventOfCode2020/types"
	"adventOfCode2020/util"
	"fmt"
)

// Run day nine of advent
func Run(input <-chan string) {
	preambleLength := 25
	xmas := util.StringChannelToIntSlice(input)
	invalidNumber := findInvalidNumber(types.NewRollingCypher(xmas[:preambleLength]), xmas[preambleLength:])
	fmt.Println("Day9 -- Part1: ", invalidNumber)
	fmt.Println("Day9 -- Part2: ", encryptionWeakness(xmas, invalidNumber))
}

func findInvalidNumber(cypher *types.RollingCypher, input []int) int {
	for _, i := range input {
		if !cypher.Validate(i) {
			return i
		}
		cypher.Increment(i)
	}
	panic("Never found the invalid number!")
}

func encryptionWeakness(input []int, invalidNumber int) int {
	tail, head := 0, 1
	sum := input[tail] + input[head]
	for tail < head && (head < len(input)-1) {
		if sum < invalidNumber {
			head++
			sum += input[head]
		}
		if sum > invalidNumber {
			sum -= input[tail]
			tail++
		}
		if sum == invalidNumber {
			min, max := input[tail], input[tail]
			for i := tail; i <= head; i++ {
				if input[i] > max {
					max = input[i]
				}
				if input[i] < min {
					min = input[i]
				}
			}
			return min + max
		}
	}
	panic("Never found encryption weakness!")
}
