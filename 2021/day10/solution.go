package day10

import (
	"bufio"
	"bytes"
	"sort"
)

func Part1(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	var total int
	for scanner.Scan() {
		l := scanner.Text()
		illegal := findIllegalChar(l)

		if illegal != "" {
			total += points[string(illegal)]
		}
	}

	return total
}

func Part2(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	points := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}

	var linenum int
	var totals []int
	for scanner.Scan() {
		l := scanner.Text()
		missing := findMissingChars(l)

		if len(missing) > 0 {
			var total int
			for _, m := range missing {
				total *= 5
				total += points[m]
			}
			totals = append(totals, total)
		}
		linenum++
	}

	sort.Ints(totals)
	idx := len(totals) / 2

	return totals[idx]
}

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

func findMissingChars(line string) []string {
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
			return nil
		}
		unmatchedOpen = unmatchedOpen[0 : len(unmatchedOpen)-1]
	}

	var missing []string
	for _, char := range unmatchedOpen {
		missing = append(missing, pairs[char])
	}

	reverse(missing)

	return missing
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}
