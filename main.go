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
	day1()
}

func getMass(number int) int {
	return number/3 - 2
}

func day1() {
	file, err := os.Open("input/01.txt")
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	var sum int
	var line string

	for {
		line, err = reader.ReadString('\n')
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
