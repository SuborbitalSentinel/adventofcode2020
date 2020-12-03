package day2

import (
	"adventOfCode2020/util"
	"fmt"
	"strconv"
)

// Run day two of advent
func Run(input <-chan string) {
	passwordList := util.ChannelToSlice(input)

	part1, part2 := validatePasswords(passwordList)
	fmt.Println("Day2 -- Part1: ", part1)
	fmt.Println("Day2 -- Part2: ", part2)
}

func validatePasswords(input []string) (part1, part2 int) {
	regex := `(?P<Min>\d+)-(?P<Max>\d+) (?P<Search>\w): (?P<Password>\w+)`
	for _, i := range input {
		m := util.RegexMap(regex, i)
		password := m["Password"]
		searchLetter := []rune(m["Search"])[0]
		min, _ := strconv.Atoi(m["Min"])
		max, _ := strconv.Atoi(m["Max"])

		if validateOccurance(password, searchLetter, min, max) {
			part1++
		}
		if validatePosition(password, searchLetter, min, max) {
			part2++
		}
	}
	return
}

func validatePosition(password string, letter rune, pos1, pos2 int) bool {
	in1 := []rune(password)[pos1-1] == letter
	in2 := []rune(password)[pos2-1] == letter

	return (in1 || in2) && !(in1 && in2)
}

func validateOccurance(password string, letter rune, min, max int) bool {
	m := make(map[rune]int)
	for _, char := range password {
		if _, ok := m[char]; !ok {
			m[char] = 1
		} else {
			m[char]++
		}
	}
	return min <= m[letter] && m[letter] <= max
}
