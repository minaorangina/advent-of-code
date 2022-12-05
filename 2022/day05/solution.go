package day05

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func Part1(b []byte) string {
	scanner := bufio.NewScanner(bytes.NewReader(b))

	split := regexp.MustCompile(`^(\s+\d+\s+)+$`)

	outerIdx := -1

	var matrix [][]string
	var instructions []string
	var endOfCrates bool

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if split.MatchString(line) {
			endOfCrates = true
			continue
		}

		if endOfCrates {
			instructions = append(instructions, line)
			continue
		}

		outerIdx++

		matrix = append(matrix, []string{})
		asRunes := []rune(line)

		for i := 1; i < len(asRunes); i = i + 4 {
			toAppend := string(asRunes[i])
			if toAppend == " " {
				toAppend = ""
			}
			matrix[outerIdx] = append(matrix[outerIdx], toAppend)
		}
	}

	width := len(matrix[0])

	for _, ins := range instructions {
		fmt.Println(matrix, ins)

		parts := strings.Split(ins, " ")

		numCrates, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		srcCol, err := strconv.Atoi(parts[3])
		if err != nil {
			log.Fatal(err)
		}
		dstCol, err := strconv.Atoi(parts[5])
		if err != nil {
			log.Fatal(err)
		}

		srcCol, dstCol = srcCol-1, dstCol-1

		for i := 0; i < numCrates; i++ {
			var dstRow int

			// let's get the coords for where we're going in advance.
			// peek first element of the dstCol. make more space if necessary
			if matrix[0][dstCol] != "" {
				fmt.Println("before new matrix", matrix)
				newRow := make([]string, width)
				matrix = append([][]string{newRow}, matrix...)
				fmt.Println("new matrix", matrix)
			}
			// find first empty space in dstCol
		destination:
			for row := 0; row < len(matrix); row++ {
				if matrix[row][dstCol] != "" {
					break destination
				}
				dstRow = row
			}

			fmt.Println("before move", matrix)

			// pluck topmost crate from srcCol
		moveCrate:
			for row := 0; row < len(matrix); row++ {
				if matrix[row][srcCol] == "" {
					continue
				}

				crate := matrix[row][srcCol]
				matrix[row][srcCol] = ""
				fmt.Printf("Crate %s (%d) will move to [%d,%d]\n", crate, []rune(crate), dstRow, dstCol)

				matrix[dstRow][dstCol] = crate

				fmt.Println("after move", matrix)
				fmt.Println("-----\nblank line\n------")
				break moveCrate
			}

		}

	}

	answer := make([]string, width)

	for col := 0; col < width; col++ {
		for row := 0; row < len(matrix); row++ {
			nabbed := matrix[row][col]
			fmt.Println(nabbed)
			if nabbed != "" {
				answer[col] = nabbed
				break
			}
		}
	}

	return strings.Join(answer[:], "")
}
