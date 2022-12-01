package day01

import (
	"bufio"
	"io"
	"strconv"
)

func Part1(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	var max, sum int

	for scanner.Scan() {
		line := scanner.Text()

		cals, err := strconv.Atoi(line)
		if err != nil {
			if sum > max {
				max = sum
			}
			sum = 0
			continue
		}
		sum += cals
	}

	return max
}
