package data

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Parse() []string {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(bytes), "\n")
}

func ToIntSlice(input []string) []int {
	var output []int

	for _, v := range input {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, num)
	}

	return output
}
