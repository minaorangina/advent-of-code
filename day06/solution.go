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
