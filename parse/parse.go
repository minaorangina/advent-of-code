package parse

import (
	"log"
	"strconv"
)

func ToIntSlice(input []byte) []int {
	var output []int

	var cursor int
	for i, v := range input {
		if v != '\n' {
			continue
		}

		numAsBytes := input[cursor : i-1]
		cursor = i + 1

		num, err := strconv.Atoi(string(numAsBytes))
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, num)
	}

	return output
}
