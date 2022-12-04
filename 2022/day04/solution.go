package day04

import (
	"2022/utils"
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Part1(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var count int

	for scanner.Scan() {
		raw := scanner.Text()
		var values []int

		temp := strings.Split(raw, ",")
		for _, t := range temp {
			nums, _ := utils.MapToInt(strings.Split(t, "-"))
			values = append(values, nums...)
		}
		fmt.Println(values)

		elfA, elfB := values[0:2], values[2:4]
		// identify largest range
		// first elf has largest range
		elfARange, elfBRange := (elfA[1] - elfA[0]), (elfB[1] - elfB[0])

		if elfARange < elfBRange {
			if elfB[0] <= elfA[0] && elfB[1] >= elfA[1] {
				fmt.Printf("case 2: %d - %d contains %d - %d\n", elfB[0], elfB[1], elfA[0], elfA[1])
				count++
			}
		} else if elfBRange < elfARange {
			if elfA[0] <= elfB[0] && elfA[1] >= elfB[1] {
				fmt.Printf("case 1: %d - %d contains %d - %d\n", elfA[0], elfA[1], elfB[0], elfB[1])
				count++
			}
		} else if elfARange == elfBRange {
			if elfA[0] == elfB[0] && elfA[1] == elfB[1] {
				count++
			}
		}
	}

	return count
}

func Part2(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var count int

	for scanner.Scan() {
		raw := scanner.Text()
		temp := strings.Split(raw, ",")
		var values []int

		for _, t := range temp {
			nums, _ := utils.MapToInt(strings.Split(t, "-"))
			values = append(values, nums...)
		}

		elfA, elfB := values[0:2], values[2:4]

		// find elf with the lowest starting number
		if elfA[0] < elfB[0] {
			if elfA[1] >= elfB[0] {
				count++
			}
		} else if elfB[0] < elfA[0] {
			if elfB[1] >= elfA[0] {
				count++
			}
		} else {
			// this must mean the starting numbers for the elves are the same
			count++
		}

	}

	return count
}
