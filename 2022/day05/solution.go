package day05

import (
	"bufio"
	"bytes"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type direction int

const (
	fromTop direction = iota
	fromBottom
)

func Solution(b []byte, plonkDirection direction) string {
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
		// fmt.Println(matrix, ins)

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

		// for each crate, collect in a queue
		toMove := make([]string, numCrates)

		// fmt.Printf("looking for %d crates\n", numCrates)

		for i := 0; i < numCrates; i++ {
			for row := 0; row < len(matrix); row++ {
				// fmt.Println("locating crate:", matrix[row][srcCol])
				if matrix[row][srcCol] == "" {
					continue
				}
				// fmt.Println("continuing with crate:", matrix[row][srcCol])
				// found a crate, put in the queue
				idx := i
				if plonkDirection == fromTop {
					idx = numCrates - i - 1
				}
				toMove[idx] = pluckCrate(matrix, row, srcCol)
				break
			}
		}

		// fmt.Println("to move", toMove)

		// peek for empty space(s)
		spacesReq := numSpacesRequired(matrix, dstCol, numCrates)
		// fmt.Println("num empty spaces req", spacesReq)

		for i := 0; i < spacesReq; i++ {
			// fmt.Println("before new matrix", matrix)
			newRow := make([]string, width)
			matrix = append([][]string{newRow}, matrix...)
			// fmt.Println("new matrix", matrix)
		}

		// move to destination
		for _, crate := range toMove {
			for dstRow := range matrix {
				// if this space is available
				if matrix[dstRow][dstCol] == "" {
					// and if the next one isnt, or we're at the end
					if dstRow == len(matrix)-1 || matrix[dstRow+1][dstCol] != "" {

						// fmt.Printf("about to move %s to [%d,%d] \n", crate, dstRow, dstCol)
						matrix[dstRow][dstCol] = crate
						break
					}
				}
			}
		}

		// fmt.Println("after move", matrix)
		// fmt.Println("-----\nblank line\n------")
	}

	answer := make([]string, width)

	for col := 0; col < width; col++ {
		for row := 0; row < len(matrix); row++ {
			nabbed := matrix[row][col]
			if nabbed != "" {
				answer[col] = nabbed
				break
			}
		}
	}

	return strings.Join(answer[:], "")
}

func pluckCrate(matrix [][]string, row, srcCol int) string {
	crate := matrix[row][srcCol]
	matrix[row][srcCol] = ""
	return crate
}

func numSpacesRequired(matrix [][]string, dstCol, numCrates int) int {
	spacesReq := 0

	for i := 0; i < numCrates; i++ {
		if matrix[i][dstCol] != "" { // no space at the inn
			spacesReq++
		}
	}

	return spacesReq
}
