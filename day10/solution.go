package day10

import (
	"bufio"
	"bytes"
)

func Part1(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	var total, linenum int
	for scanner.Scan() {
		l := scanner.Text()
		illegal := findIllegalChar(l)

		if illegal != "" {
			total += points[string(illegal)]
		}
		linenum++
	}

	return total
}

/*
corrupted
when you encounter a closing char that you were not expecting
- keep track of the last seen open character on a stack, pop the stack when a pair is found.

incomplete
when you reach the end of the line and still have open characters remaining
- count the open char stack at the end.

1. identify the corrupted lines
expecting 2, 4, 5, 7, 8 (0-idx)
*/

func findIllegalChar(line string) string {
	pairs := map[string]string{
		"(": ")", "{": "}", "[": "]", "<": ">",
	}

	var unmatchedOpen []string

	for _, c := range line {
		char := string(c)
		if _, ok := pairs[char]; ok {
			unmatchedOpen = append(unmatchedOpen, char)
			continue
		}

		lastSeenOpen := unmatchedOpen[len(unmatchedOpen)-1]
		expectedClosed := pairs[lastSeenOpen]

		if char != expectedClosed {
			// illegal
			return char
		}
		unmatchedOpen = unmatchedOpen[0 : len(unmatchedOpen)-1]
	}

	return ""
}
