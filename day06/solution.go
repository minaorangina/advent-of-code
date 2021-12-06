package day06

import (
	"log"
	"strconv"
	"strings"
)

func Part1(inputB []byte, days int) int {
	asString := strings.Split((string(inputB)), ",")

	var input []int
	for _, c := range asString {
		n, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, n)
	}

	for d := 0; d < days; d++ {
		for i := range input {
			if input[i] == 0 {
				input[i] = 6
				input = append(input, 8)
				continue
			}
			input[i]--
		}
	}
	return len(input)
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

	var results []int
	for _, v := range input {
		results = append(results, doStuff(v)...)
	}
	return len(results)
}

func doStuff(n int) []int {
	return nil
}

/*
Rules
- each fish <7 has the same number after 7 days
- each new fish will have 1 day less than its parent

- each fish  >=7 has self-7 after 7 days

q - give me a new number
q - give me a new fish (if i should have one	)
*/
