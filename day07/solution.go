package day07

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func Part1(inputB []byte) int {
	asString := strings.Split((string(inputB)), ",")

	var input []int
	for _, c := range asString {
		n, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, n)
	}

	cost := map[int]struct{}{}
	for _, v := range input {
		cost[v] = struct{}{}
	}

	smallest := math.Inf(1)
	for position := range cost {
		var c float64
		for _, v := range input {
			diff := math.Abs(float64(v - position))
			c += diff
		}
		if c < smallest {
			smallest = c
		}
	}

	return int(smallest)
}

func Part2(inputB []byte) int {
	asString := strings.Split((string(inputB)), ",")

	var input []int
	for _, c := range asString {
		n, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, n)
	}

	return 0
}
