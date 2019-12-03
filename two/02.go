package two

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/minaorangina/advent-of-code-2019/utils"
)

func Exec() {
	fmt.Println("Day 2")
	part1()
}

func part1() {
	type Op int

	const (
		Add      Op = 1
		Multiply Op = 2
		Exit     Op = 99
	)

	values := getValues()

	values[1] = 12
	values[2] = 2

	for i := 0; i < len(values); i += 4 {
		switch Op(values[i]) {
		case Add:
			values[values[i+3]] = values[values[i+1]] + values[values[i+2]]
		case Multiply:
			values[values[i+3]] = values[values[i+1]] * values[values[i+2]]
		case Exit:
			break
		default:
			// fmt.Println("Unknown op code")
			break
		}
	}

	fmt.Println("	Part 1:", values[0])
}

func getValues() []int {
	fileBytes, err := ioutil.ReadFile("input/02.txt")
	utils.Check(err)
	file := string(fileBytes)

	valuesStr := strings.Split(
		strings.TrimRight(string(file), "\n"),
		",",
	)

	values := []int{}

	for _, val := range valuesStr {
		num, err := strconv.Atoi(val)
		utils.Check(err)
		values = append(values, num)
	}
	return values
}
