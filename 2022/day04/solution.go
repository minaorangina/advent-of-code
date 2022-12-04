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

		// identify largest range
		// first elf has largest range
		firstElfRange, secondElfRange := (values[1] - values[0]), (values[3] - values[2])

		if firstElfRange > secondElfRange {
			if values[0] <= values[2] && values[1] >= values[3] {
				fmt.Printf("case 1: %d - %d contains %d - %d\n", values[0], values[1], values[2], values[3])
				count++
			}
		} else if secondElfRange > firstElfRange {
			if values[2] <= values[0] && values[3] >= values[1] {
				fmt.Printf("case 2: %d - %d contains %d - %d\n", values[2], values[3], values[0], values[1])
				count++
			}
		} else if firstElfRange == secondElfRange {
			if values[0] == values[2] && values[1] == values[3] {
				count++
			}
		}
	}

	return count
}
