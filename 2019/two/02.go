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
	fmt.Println("Part 1:", part1())
	fmt.Println("Part 2:", part2())
}

func part1() int {

	values := getValues()

	// inputs
	values[1] = 12 // noun
	values[2] = 2  // verb

	return runIntCode(values)
}

func part2() int {
	const max = 99
	const expected = 19690720

	values := getValues()

	for noun := 1; noun < 100; noun++ {
		for verb := 1; verb < 100; verb++ {

			copiedValues := make([]int, len(values))
			copy(copiedValues, values)
			copiedValues[1] = noun
			copiedValues[2] = verb

			result := runIntCode(copiedValues)

			if result == expected {
				return 100*noun + verb
			}
		}
	}
	panic("Failed")
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

func runIntCode(values []int) int {
	type Op int

	const (
		Add      Op = 1
		Multiply Op = 2
		Exit     Op = 99
	)

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

	return values[0]
}
