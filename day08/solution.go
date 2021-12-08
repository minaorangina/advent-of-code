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
		knownLengths := map[int]int{
			2: 1,
			3: 7,
			4: 4,
			7: 8,
		}

		// done when len == 9
		encodingToDigit := map[string]int{}
		digitToEncoding := map[int]string{}

		positions := map[string]string{
			"top":         "",
			"upperleft":   "",
			"upperright":  "",
			"middle":      "",
			"bottomleft":  "",
			"bottomright": "",
			"bottom":      "",
		}

		possiblePositions := map[string][]string{
			"top":         {"a", "b", "c", "d", "e", "f", "g"},
			"upperleft":   {"a", "b", "c", "d", "e", "f", "g"},
			"upperright":  {"a", "b", "c", "d", "e", "f", "g"},
			"middle":      {"a", "b", "c", "d", "e", "f", "g"},
			"bottomleft":  {"a", "b", "c", "d", "e", "f", "g"},
			"bottomright": {"a", "b", "c", "d", "e", "f", "g"},
			"bottom":      {"a", "b", "c", "d", "e", "f", "g"},
		}

		_ = positions
		_ = possiblePositions
		_ = outputVals

		for _, encoding := range encodingRules {
			// identify 1, 4, 7 and 8
			digit, ok := knownLengths[len(encoding)]
			if ok {
				canon := canonicalEncoding(encoding)
				encodingToDigit[canon] = digit
				digitToEncoding[digit] = canon
			}
		}

		//

		fmt.Println("iudgfiusdh", encodingToDigit)
	}

	return count
}

func canonicalEncoding(raw string) string {
	canon := []rune(raw)
	sort.Slice(canon, func(a, b int) bool {
		return canon[a] < canon[b]
	})
	return string(canon)
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
