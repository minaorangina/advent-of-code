package day05

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Part1(input []byte) int {
	points := map[string]int{}

	scanner := bufio.NewScanner(strings.NewReader(string(input)))

	for scanner.Scan() {
		l := scanner.Text()
		coordsS := strings.Split(l, " -> ")

		var coords [][]int
		for _, co := range coordsS {
			spl := strings.Split(co, ",")
			var nums []int
			for _, char := range spl {
				n, err := strconv.Atoi(char)
				if err != nil {
					log.Fatal(err)
				}
				nums = append(nums, n)
			}
			coords = append(coords, nums)
		}

		if !horizontalOrVertical(coords[0], coords[1]) {
			continue
		}

		nextCoords := coords[0]
		for {
			// add to map
			points[fmt.Sprintf("%d,%d", nextCoords[0], nextCoords[1])]++
			if coordsEqual(nextCoords, coords[1]) {
				break
			}

			nextCoords = doTing(nextCoords, coords[1])
		}
	}

	var twoOrMore int
	for _, v := range points {
		if v >= 2 {
			twoOrMore++
		}
	}

	return twoOrMore
}

func Part2(input []byte) int {
	points := map[string]int{}

	scanner := bufio.NewScanner(strings.NewReader(string(input)))

	for scanner.Scan() {
		l := scanner.Text()
		coordsS := strings.Split(l, " -> ")

		var coords [][]int
		for _, co := range coordsS {
			spl := strings.Split(co, ",")
			var nums []int
			for _, char := range spl {
				n, err := strconv.Atoi(char)
				if err != nil {
					log.Fatal(err)
				}
				nums = append(nums, n)
			}
			coords = append(coords, nums)
		}

		nextCoords := coords[0]
		for {
			// add to map
			points[fmt.Sprintf("%d,%d", nextCoords[0], nextCoords[1])]++
			if coordsEqual(nextCoords, coords[1]) {
				break
			}

			nextCoords = doTing(nextCoords, coords[1])
		}
	}

	var twoOrMore int
	for _, v := range points {
		if v >= 2 {
			twoOrMore++
		}
	}

	return twoOrMore
}

func horizontalOrVertical(a, b []int) bool {
	return a[0] == b[0] || a[1] == b[1]
}

func coordsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func doTing(start, end []int) []int {
	// look at the x value first, determine if its the same or different
	// produce next value, track, repeat until end value is produced (track it)
	nextX, nextY := start[0], start[1]
	if start[0] != end[0] {
		// figure out nextX
		if start[0] < end[0] {
			nextX++
		} else {
			nextX--
		}
	}
	if start[1] != end[1] {
		// figure out nextY
		if start[1] < end[1] {
			nextY++
		} else {
			nextY--
		}
	}

	return []int{nextX, nextY}
}

/*
how many points have at least two overlapping lines?

- could list all existing points and count

how to determine horizontal/vertical lines

*/
