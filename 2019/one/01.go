package one

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/minaorangina/advent-of-code-2019/utils"
)

func Exec() {
	fmt.Println("Day 1")

	file, err := os.Open("input/01.txt")
	utils.Check(err)
	defer file.Close()

	part1(file)
	part2(file)
}

func getMass(number int) int {
	return number/3 - 2
}
func getMassRecursive(number int) int {
	fuel := number/3 - 2
	if fuel <= 0 {
		return 0
	}
	return fuel + getMassRecursive(fuel)
}

func part1(file *os.File) {
	reader := bufio.NewReader(file)
	var sum int

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		trimmedLine := strings.TrimSuffix(line, "\n")
		num, err := strconv.Atoi(trimmedLine)
		utils.Check(err)

		sum += getMass(num)
	}
	fmt.Println("	Part 1:", sum)
}

func part2(file *os.File) {
	// keep getting zero if this isn't here...
	file, err := os.Open("input/01.txt")
	utils.Check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	var sum int

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		trimmedLine := strings.TrimSuffix(line, "\n")
		num, err := strconv.Atoi(trimmedLine)
		utils.Check(err)

		sum += getMassRecursive(num)
	}
	fmt.Println("	Part 2:", sum)
}
