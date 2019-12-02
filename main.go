package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	file, err := os.Open("input/01.txt")
	check(err)
	defer file.Close()

	day1(file)
	day2(file)
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

func day1(file *os.File) {
	reader := bufio.NewReader(file)
	var sum int

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		trimmedLine := strings.TrimSuffix(line, "\n")
		num, err := strconv.Atoi(trimmedLine)
		check(err)

		sum += getMass(num)
	}
	fmt.Println("Day 01:", sum)
}

func day2(file *os.File) {
	// keep getting zero if this isn't here...
	file, err := os.Open("input/01.txt")
	check(err)
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
		check(err)

		sum += getMassRecursive(num)
	}
	fmt.Println("Day 02:", sum)
}
