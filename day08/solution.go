package day08

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

func Part1(input []byte) int {

	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	var count int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		outputVals := strings.Fields(parts[1])

		for _, digit := range outputVals {
			if uniqueLength(len(digit)) {
				count++
			}
		}
	}

	segmentCounts := map[int]int{
		1: 2,
		4: 4,
		7: 3,
		8: 7,
	}
	_ = segmentCounts
	return count
}

func Part2(input []byte) int {

	scanner := bufio.NewScanner(strings.NewReader(string(input)))
	var count int
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")

		encodingRules := strings.Fields(parts[0])
		outputVals := strings.Fields(parts[1])

		// segmentCount: digit
		knownEncodings := map[int]int{
			2: 1,
			3: 7,
			4: 4,
			7: 8,
		}

		encodingToDigit := map[string]int{}

		positions := map[string]*int{
			"top":         nil,
			"upperleft":   nil,
			"upperright":  nil,
			"middle":      nil,
			"bottomleft":  nil,
			"bottomright": nil,
			"bottom":      nil,
		}

		_ = positions
		_ = outputVals

		for _, encoding := range encodingRules {
			// identify 1 (2 chars, upper right, lower right, no ordering)
			digit, ok := knownEncodings[len(encoding)]
			if ok {
				canonicalEncoding := []rune(encoding)
				sort.Slice(canonicalEncoding, func(a, b int) bool {
					return canonicalEncoding[a] < canonicalEncoding[b]
				})

				encodingToDigit[string(canonicalEncoding)] = digit
			}

			// identify 7 -> find spare char (top)

			// identify 4 -> find spare 2 chars (middle, upper left, no ordering)

			// identify 8 -> find spare 2 chars (bottom, lower left, no ordering)

			// unknowns

			// identify 9 -> 4 + top char + find spare (bottom)

		}

		fmt.Println("iudgfiusdh", encodingToDigit)
	}

	return count
}

func uniqueLength(l int) bool {
	switch l {
	case 2, 3, 4, 7:
		return true
	}
	return false
}

// 7 (1 plus 1xline)
// 4 (1 plus 2xline)
// 8 (all lines)
