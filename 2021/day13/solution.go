package day13

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
	"strings"
)

func Part1(input []byte) int {
	parts := bytes.Split(input, []byte("\n\n"))

	coords := map[[2]int]struct{}{}
	folds := map[string]int{}

	scanner := bufio.NewScanner(bytes.NewReader(parts[0]))
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ",")

		var coord [2]int
		for i, s := range nums {
			num, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			coord[i] = num
		}
		coords[coord] = struct{}{}
	}

	instructions := bytes.Split(parts[1], []byte("\n"))
	var foldAxis string

	for _, f := range instructions[0:1] {
		h := bytes.Replace(f, []byte("fold along "), []byte(""), 1)
		foldAxis = string(h[0])

		num, err := strconv.Atoi(string(h[2:]))
		if err != nil {
			log.Fatal(err)
		}

		folds[foldAxis] = num
	}

	var idx int
	if foldAxis == "y" {
		idx = 1
	}

	for coord := range coords {
		target, linePos := coord[idx], folds[foldAxis]
		if target > linePos {
			delete(coords, coord)

			diff := target - linePos
			newPos := linePos - diff

			coord[idx] = newPos
			coords[coord] = struct{}{}
		}
	}

	return len(coords)
}
