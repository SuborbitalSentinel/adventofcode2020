package day4

import (
	"adventOfCode2020/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func parsePassportList(input []string) (passports []passport) {
	currentPassport := newPassport()
	for _, line := range input {
		if len(line) == 0 {
			passports = append(passports, currentPassport)
			currentPassport = newPassport()
		} else {
			fields := strings.Split(line, " ")
			for _, field := range fields {
				pair := strings.Split(field, ":")
				currentPassport.Fields[pair[0]] = pair[1]
			}
		}
	}
	passports = append(passports, currentPassport)
	return
}

// Run day four of advent
func Run(input <-chan string) {
	passports := parsePassportList(util.ChannelToSlice(input))

	fmt.Println("Day4 -- Part1: ", part1(passports))
	fmt.Println("Day4 -- Part2: ", part2(passports))
}

func part1(input []passport) (count int) {
	for _, p := range input {
		if p.ContainsKeys(requiredFields...) {
			count++
		}
	}
	return
}

func part2(input []passport) (count int) {
	for _, p := range input {
		if p.ContainsKeys(requiredFields...) && p.IsValid() {
			count++
		}
	}
	return
}

type passport struct {
	Fields map[string]string
}

func newPassport() passport {
	return passport{make(map[string]string)}
}

func (p passport) IsValid() bool {
	if value, ok := p.Fields["byr"]; ok && !validBirthYear(value) {
		return false
	}
	if value, ok := p.Fields["iyr"]; ok && !validIssueYear(value) {
		return false
	}
	if value, ok := p.Fields["eyr"]; ok && !validExpirationYear(value) {
		return false
	}
	if value, ok := p.Fields["hgt"]; ok && !validHeight(value) {
		return false
	}
	if value, ok := p.Fields["hcl"]; ok && !validHairColor(value) {
		return false
	}
	if value, ok := p.Fields["ecl"]; ok && !validEyeColor(value) {
		return false
	}
	if value, ok := p.Fields["pid"]; ok && !validPassportID(value) {
		return false
	}
	return true
}

func (p passport) ContainsKeys(key ...string) bool {
	for _, k := range key {
		if _, ok := p.Fields[k]; !ok {
			return false
		}
	}
	return true
}

func validBirthYear(birthYear string) bool {
	by, _ := strconv.Atoi(birthYear)
	return 1920 <= by && by <= 2002
}

func validIssueYear(issueYear string) bool {
	iy, _ := strconv.Atoi(issueYear)
	return 2010 <= iy && iy <= 2020
}

func validExpirationYear(expirationYear string) bool {
	ey, _ := strconv.Atoi(expirationYear)
	return 2020 <= ey && ey <= 2030
}

func validHeight(height string) bool {
	cmrgx := regexp.MustCompile(`^\d+cm$`)
	inrgx := regexp.MustCompile(`^\d+in$`)

	if cmrgx.MatchString(height) {
		h, _ := strconv.Atoi(height[:len(height)-2])
		return 150 <= h && h <= 193
	}

	if inrgx.MatchString(height) {
		h, _ := strconv.Atoi(height[:len(height)-2])
		return 59 <= h && h <= 76
	}
	return false
}

func validHairColor(hairColor string) bool {
	r := regexp.MustCompile(`^#[a-f0-9]{6}$`)
	return r.MatchString(hairColor)
}

func validEyeColor(eyeColor string) bool {
	r := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	return r.MatchString(eyeColor)
}

func validPassportID(passportID string) bool {
	r := regexp.MustCompile(`^\d{9}$`)
	return r.MatchString(passportID)
}
