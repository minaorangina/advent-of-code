package day03

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

func Part1(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var priorities int
	_ = priorities

	for scanner.Scan() {
		raw := scanner.Text()

		halfway := len(raw) / 2
		comp1 := raw[0:halfway]
		comp2 := raw[halfway:]

		var duplicate rune

		for _, char := range comp2 {
			if strings.ContainsRune(comp1, char) {
				duplicate = char
				break
			}
		}

		priority := getPriority(duplicate)

		fmt.Println("duplicate", string(duplicate), "ASCII", duplicate, "priority", priority)
		priorities += priority
	}

	return priorities
}

func Part2(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var priorities int
	lines := []string{}

	for scanner.Scan() {
		raw := scanner.Text()

		lines = append(lines, raw)
		if len(lines) < 3 {
			continue
		}

		counts := map[rune]int{}

		for _, line := range lines {
			seen := map[rune]struct{}{}

			for _, char := range line {
				if _, ok := seen[char]; ok {
					continue
				}

				counts[char] = counts[char] + 1
				seen[char] = struct{}{}
			}
		}

		var badge rune

		for char, v := range counts {
			if v == 3 {
				badge = char
			}
		}

		priority := getPriority(badge)

		fmt.Println("badge", string(badge), "ASCII", badge, "priority", priority)
		priorities += priority

		lines = []string{}
	}

	return priorities
}

func getPriority(c rune) int {
	var priority int
	if unicode.IsLower(c) {
		priority = int(c - 'a' + 1)
	} else {
		priority = int(c - 'A' + 27)
	}
	return priority
}
