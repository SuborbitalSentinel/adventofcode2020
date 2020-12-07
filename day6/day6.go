package day6

import (
	"adventOfCode2020/types"
	"adventOfCode2020/util"
	"fmt"
)

func parseGroups(input []string) (groups []group) {
	currentGroup := group{Answers: make(types.StringSet)}
	for _, line := range input {
		if len(line) == 0 {
			groups = append(groups, currentGroup)
			currentGroup = group{Answers: make(types.StringSet)}
		} else {
			currentGroup.Participants++
			for _, c := range line {
				currentGroup.Answers.Add(string(c))
			}
		}
	}
	groups = append(groups, currentGroup)
	return
}

// Run day six of advent
func Run(input <-chan string) {
	groups := parseGroups(util.ChannelToSlice(input))

	totalAnswers := 0
	totalUnanamousAnswers := 0
	for _, group := range groups {
		totalAnswers += len(group.Answers)
		totalUnanamousAnswers += len(group.Answers.SeenExactly(group.Participants))
	}

	fmt.Println("Day6 -- Part1: ", totalAnswers)
	fmt.Println("Day6 -- Part2: ", totalUnanamousAnswers)
}

type group struct {
	Participants int
	Answers      types.StringSet
}
