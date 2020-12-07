package day7

import (
	"adventOfCode2020/util"
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input []string) (bags []bag) {
	for _, line := range input {
		key, values := splitIntoKeysAndValues(line)
		g := bag{Key: key, Values: make(map[string]int)}
		for _, value := range values {
			g.AddValue(value)
		}
		bags = append(bags, g)
	}
	return
}

// Run day seven of advent
func Run(input <-chan string) {
	bags := parseInput(util.ChannelToSlice(input))
	bagMap := make(map[string]bag)
	for _, bag := range bags {
		bagMap[bag.Key] = bag
	}

	fmt.Println("Day7 -- Part1: ", part1(bags, bagMap))
	fmt.Println("Day7 -- Part2: ", bagMap["shiny gold bag"].BagCount(bagMap)-1)
}

func part1(bags []bag, m map[string]bag) int {
	myBag := "shiny gold bag"
	containsMyBag := 0
	for _, bag := range bags {
		if bag.ContainsBag(myBag, m) {
			containsMyBag++
		}
	}
	return containsMyBag
}

func splitIntoKeysAndValues(line string) (key string, values []string) {
	l := strings.Split(line, "contain")
	key = cleanKey(l[0])
	for _, value := range strings.Split(l[1], ",") {
		cleaned := cleanValue(value)
		if cleaned != "no other bag" {
			values = append(values, cleaned)
		}
	}
	return
}

func cleanKey(key string) string {
	k := strings.TrimSpace(key)
	return k[:len(k)-1]
}

func cleanValue(value string) string {
	v := strings.TrimSpace(strings.ReplaceAll(value, ".", ""))
	if v[len(v)-1] == 's' {
		return v[:len(v)-1]
	}
	return v
}

type bag struct {
	Key    string
	Values map[string]int
}

//This count includes the bag you are calling it on
func (g bag) BagCount(m map[string]bag) int {
	bagCount := 1
	for key, value := range g.Values {
		bagCount += value * m[key].BagCount(m)
	}
	return bagCount
}

func (g bag) ContainsBag(input string, m map[string]bag) bool {
	//Check if I contain the bag
	if _, ok := g.Values[input]; ok {
		return true
	}
	//Check if my bags contain the bag
	for key := range g.Values {
		if m[key].ContainsBag(input, m) {
			return true
		}
	}
	// I do not contain the bag
	return false
}

func (g bag) AddValue(input string) {
	r := util.RegexMap(`(?P<Count>\d+) (?P<Bag>.*)$`, input)
	b := r["Bag"]
	count, _ := strconv.Atoi(r["Count"])
	g.Values[b] = count
}
