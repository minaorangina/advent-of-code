package day07

import (
	"math"
	"sort"
)

func Part1(input []int) int {
	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j]
	})

	max := input[0]

	smallestCost := math.Inf(1)
	for position := 0; position <= max; position++ {
		var c float64
		for _, v := range input {
			diff := math.Abs(float64(v - position))
			c += diff
		}
		if c < smallestCost {
			smallestCost = c
		}
	}

	return int(smallestCost)
}

func Part2(input []int) int {
	sort.Slice(input, func(i, j int) bool {
		return input[i] > input[j]
	})

	max := input[0]

	smallestCost := math.Inf(1)
	for position := 0; position <= max; position++ {
		var c float64
		for _, v := range input {
			diff := math.Abs(float64(v - position))
			c += getScore(diff)
		}
		if c < smallestCost {
			smallestCost = c
		}
	}

	return int(smallestCost)
}

func getScore(diff float64) float64 {
	var score int
	for i := 1; i <= int(diff); i++ {
		score += i
	}
	return float64((score))
}
