package day01

import (
	"2022/utils"
	"bufio"
	"io"
	"sort"
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

func Part2(input io.Reader) int {
	scanner := bufio.NewScanner(input)

	var (
		sum      int
		calories []int
	)

	for scanner.Scan() {
		line := scanner.Text()

		cals, err := strconv.Atoi(line)
		if err != nil {
			if sum > 0 {
				calories = append(calories, sum)
				sum = 0
			}
			continue
		}
		sum += cals
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	return utils.Reduce(calories[0:3], func(acc, current int) int {
		return acc + current
	}, 0)
}
