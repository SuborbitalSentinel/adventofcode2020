package util

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

// RegexMap returns a map of regex labels to values
func RegexMap(regex, input string) (m map[string]string) {
	r := regexp.MustCompile(regex)
	match := r.FindStringSubmatch(input)

	m = make(map[string]string)

	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			m[name] = match[i]
		}
	}

	return
}

// Reverse a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ChannelToSlice returns an slice insead of a channel
func ChannelToSlice(input <-chan string) []string {
	var output []string

	for input := range input {
		output = append(output, input)
	}

	return output
}

// StringChannelToIntSlice returns a int slice insead of a channel
func StringChannelToIntSlice(input <-chan string) []int {
	var output []int

	for input := range input {
		value, _ := strconv.Atoi(input)
		output = append(output, value)
	}

	return output
}

// Readwords returns a channel containing the content of the file
func Readwords(path string) <-chan string {
	fobj, err := os.Open(path)
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(fobj)
	if err := scanner.Err(); err != nil {
		return nil
	}

	scanner.Split(bufio.ScanWords)
	chnl := make(chan string)
	go func() {
		for scanner.Scan() {
			chnl <- scanner.Text()
		}
		close(chnl)
	}()

	return chnl
}

// Readints returns a channel containing the content of the file with int type
func Readints(path string) <-chan int {
	fobj, err := os.Open(path)
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(fobj)
	if err := scanner.Err(); err != nil {
		return nil
	}

	chnl := make(chan int)
	go func() {
		for scanner.Scan() {
			x, _ := strconv.Atoi(scanner.Text())
			chnl <- x
		}
		close(chnl)
	}()

	return chnl

}

// Readlines returns a channel containing the content of the file
func Readlines(path string) <-chan string {
	fobj, err := os.Open(path)
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(fobj)
	if err := scanner.Err(); err != nil {
		return nil
	}

	chnl := make(chan string)
	go func() {
		for scanner.Scan() {
			chnl <- scanner.Text()
		}
		close(chnl)
	}()

	return chnl
}

// Abs Go REALLY doesn't have a built in abs function for integers....WHAII
func Abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}
