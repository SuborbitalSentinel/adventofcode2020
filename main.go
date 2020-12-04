package main

import (
	"adventOfCode2020/day1"
	"adventOfCode2020/day2"
	"adventOfCode2020/day3"
	"adventOfCode2020/util"
)

func main() {
	day1.Run(util.Readlines("day1/input"))
	day2.Run(util.Readlines("day2/input"))
	day3.Run(util.Readlines("day3/input"))
}
