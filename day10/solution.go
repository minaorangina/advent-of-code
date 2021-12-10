package day10

import (
	"bufio"
	"bytes"
	"fmt"
)

func Part1(input []byte) int {
	var corrupted []string

	scanner := bufio.NewScanner(bytes.NewReader(input))

	var linecount int
	for scanner.Scan() {
		l := scanner.Text()
		res := isCorrupted(l)
		fmt.Println(res)
		if res {
			corrupted = append(corrupted, l)
			linecount++
		}
	}

	fmt.Println(len(corrupted), corrupted)

	return 0
}

func isCorrupted(line string) bool {
	openChars := map[string]struct{}{
		"(": {}, "[": {}, "{": {}, "<": {},
	}
	closeChars := map[rune]struct{}{
		')': {}, ']': {}, '}': {}, '>': {},
	}
	_, _ = openChars, closeChars
	// TODO
	return false
}
