package util

import (
	"bufio"
	"os"
)

// Reverse a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// ChannelToSlice returns an array insead of a channel
func ChannelToSlice(input <-chan string) []string {
	var output []string

	for input := range input {
		output = append(output, input)
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
