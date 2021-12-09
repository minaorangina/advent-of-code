package day08

import (
	"bufio"
	"fmt"
	"log"
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

		/*****/

		// done when len == 9
		encodingToDigit := map[string]int{}
		digitToEncoding := map[int]string{}

		const (
			top position = iota
			upperLeft
			upperRight
			middle
			lowerLeft
			lowerRight
			bottom
		)

		positionTracker := map[string][]position{
			"a": {0, 1, 2, 3, 4, 5, 6},
			"b": {0, 1, 2, 3, 4, 5, 6},
			"c": {0, 1, 2, 3, 4, 5, 6},
			"d": {0, 1, 2, 3, 4, 5, 6},
			"e": {0, 1, 2, 3, 4, 5, 6},
			"f": {0, 1, 2, 3, 4, 5, 6},
			"g": {0, 1, 2, 3, 4, 5, 6},
		}

		_ = positionTracker
		_ = outputVals

		// segmentCount: digit
		knownLengths := map[int]int{
			2: 1,
			3: 7,
			4: 4,
			7: 8,
		}

		for _, encoding := range encodingRules {
			// identify 1, 4, 7 and 8
			digit, ok := knownLengths[len(encoding)]
			if ok {
				canon := canonicalEncoding(encoding)
				encodingToDigit[canon] = digit
				digitToEncoding[digit] = canon
			}
		}

		// work out top from 1 and 7
		fmt.Printf("1 %s, 7 %s\n", digitToEncoding[1], digitToEncoding[7])
		diff := getSetDifference(digitToEncoding[1], digitToEncoding[7])
		if len(diff) != 1 {
			log.Fatalf("expected set difference of 1, got %d: %v", len(diff), diff)
		}
		positionTracker[string(diff[0])] = []position{top}

		// work out

		fmt.Println("positionTracker", positionTracker)

		// goal, get the correct position for top

	}

	return count
}

func getSetDifference(a, b string) string {
	mb := make(map[string]struct{}, len(b))

	for _, x := range b {
		mb[string(x)] = struct{}{}
	}
	var diff string
	for _, x := range a {
		if _, found := mb[string(x)]; !found {
			diff += string(x)
		}
	}
	return diff
}

func stringSetToSlice(set map[string]struct{}) []rune {
	out := []rune{}

	return out
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

type position int

func (p position) String() string {
	switch p {
	case 0:
		return "top"
	case 1:
		return "upper left"
	case 2:
		return "upper right"
	case 3:
		return "middle"
	case 4:
		return "bottom left"
	case 5:
		return "bottom left"
	case 6:
		return "bottom"
	}
	return ""
}
