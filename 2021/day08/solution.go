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
			"b": {0, 1, 2, 3, 4, 5, 6}, //2,5
			"c": {0, 1, 2, 3, 4, 5, 6}, //1,3
			"d": {0, 1, 2, 3, 4, 5, 6}, //0
			"e": {0, 1, 2, 3, 4, 5, 6}, //2,5
			"f": {0, 1, 2, 3, 4, 5, 6},
			"g": {0, 1, 2, 3, 4, 5, 6}, //1,3
		}

		positionToLetter := map[position]string{}

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
		positionTracker[diff[0]] = []position{top}
		positionToLetter[top] = diff[0]

		// work out

		// find 9 (combo of 4 and top)
		partial := []rune(digitToEncoding[4] + positionToLetter[top])
		sort.Slice(partial, func(a, b int) bool { return partial[a] < partial[b] })

		fmt.Println(string(partial), "????????")
	}

	return count
}

func getSetDifference(a, b string) []string {
	mb := make(map[string]struct{}, len(b))
	// a, but not b
	for _, x := range b {
		mb[string(x)] = struct{}{}
	}
	notIn := []string{}
	for _, x := range a {
		if _, found := mb[string(x)]; !found {
			notIn = append(notIn, string(x))
		}
	}

	ma := make(map[string]struct{}, len(a))
	// b, but not a
	for _, x := range a {
		ma[string(x)] = struct{}{}
	}

	for _, x := range b {
		if _, found := ma[string(x)]; !found {
			notIn = append(notIn, string(x))

		}
	}

	return notIn
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
