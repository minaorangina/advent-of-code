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
		line := input[cursor:i]
		cursor = i + 1
		num, err := strconv.Atoi(string(line))
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, num)
	}

	return output
}

func ToStringSlice(input []byte) []string {
	var output []string

	var cursor int
	for i, v := range input {
		if v != '\n' {
			continue
		}
		line := input[cursor:i]
		cursor = i + 1
		output = append(output, string(line))
	}

	return output
}
