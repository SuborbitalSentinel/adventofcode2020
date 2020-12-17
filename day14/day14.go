package day14

import (
	"adventOfCode2020/util"
	"fmt"
	"regexp"
	"strconv"
)

var maskValue = regexp.MustCompile(`mask = (?P<mask>.*)$`)
var memoryValue = regexp.MustCompile(`mem\[(?P<address>\d+)\] = (?P<value>\d+)$`)

// Run executes day 14 of advent
func Run(input <-chan string) {
	puzzle := util.ChannelToSlice(input)

	p1Total := 0
	for _, value := range part1(puzzle) {
		num, _ := strconv.ParseInt(value, 2, 0)
		p1Total += int(num)
	}
	fmt.Println("Day14 -- Part1:", p1Total)

	p2Total := 0
	for _, value := range part2(puzzle) {
		p2Total += value
	}
	fmt.Println("Day14 -- Part2:", p2Total)

}

func part1(input []string) (memory map[string]string) {
	memory = make(map[string]string)
	var currentMask string
	for _, line := range input {
		if maskValue.MatchString(line) {
			currentMask = util.Params(maskValue, line)["mask"]
		} else {
			p := util.Params(memoryValue, line)
			memory[p["address"]] = applyMaskToValue(currentMask, paddedBinary(p["value"]))
		}
	}
	return
}

func part2(input []string) (memory map[string]int) {
	memory = make(map[string]int)
	var currentMask string
	for _, line := range input {
		if maskValue.MatchString(line) {
			currentMask = util.Params(maskValue, line)["mask"]
		} else {
			p := util.Params(memoryValue, line)
			addresses := expandAddress(0, []string{applyMaskToAddress(currentMask, paddedBinary(p["address"]))})
			for _, address := range addresses {
				memory[address] = toInt(p["value"])
			}
		}
	}
	return
}

func expandAddress(bit int, collector []string) []string {
	if bit == len(collector[0]) {
		return collector
	}
	newAddresses := []string{}
	for i, value := range collector {
		if v := []rune(value); v[bit] == 'X' {
			v[bit] = '0'
			collector[i] = string(v)
			v[bit] = '1'
			newAddresses = append(newAddresses, string(v))
		}
	}
	bit++
	collector = append(collector, newAddresses...)
	return expandAddress(bit, collector)
}

func applyMaskToAddress(mask, value string) string {
	if len(mask) != len(value) {
		panic("mask and value must be same length")
	}

	output := []rune{}
	for i := range value {
		v := rune(value[i])
		if mask[i] == 'X' {
			v = 'X'
		} else if mask[i] == '1' {
			v = '1'
		}
		output = append(output, v)
	}
	return string(output)

}

func applyMaskToValue(mask, value string) string {
	if len(mask) != len(value) {
		panic("mask and value must be same length")
	}

	output := []rune{}
	for i := range value {
		v := rune(value[i])
		if mask[i] != 'X' {
			v = rune(mask[i])
		}
		output = append(output, v)
	}
	return string(output)
}

func paddedBinary(input string) string {
	v, _ := strconv.Atoi(input)
	return fmt.Sprintf("%036b", v)
}

func toInt(input string) int {
	out, _ := strconv.Atoi(input)
	return out
}
