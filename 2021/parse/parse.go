package parse

import (
	"log"
	"strconv"
	"strings"
)

func ColumnToIntSlice(input []byte) []int {
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

// ToIntSlice converts a comma-separated string of numbers into a slice of integers
func ToIntSlice(inputB []byte) []int {
	asString := strings.Split((string(inputB)), ",")

	var input []int
	for _, c := range asString {
		n, err := strconv.Atoi(c)
		if err != nil {
			log.Fatal(err)
		}
		input = append(input, n)
	}
	return input
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
