package day11

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
	"strings"
)

const gridSize = 10

type matrix [][]int
type cell [2]int

func newCell(x, y int) cell {
	return [2]int{x, y}
}

func (m matrix) String() string {
	var output string
	for _, row := range m {
		for _, v := range row {
			output += strconv.Itoa(v)
		}
		output += "\n"
	}
	return output
}

func Part1(input []byte, steps int) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	var m matrix
	for scanner.Scan() {
		l := scanner.Text()
		n := strings.Split(l, "")

		var line []int
		for _, v := range n {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			line = append(line, num)
		}

		m = append(m, line)
	}

	var flashes int

	for n := 0; n < steps; n++ {
		for i, row := range m {
			for j := range row {
				m[i][j]++
			}
		}

		var queue []cell

		for i, row := range m {
			for j := range row {
				if row[j] > 9 {
					queue = append(queue, newCell(i, j))
					m[i][j] = 0
					flashes++
				}
			}
		}

		for len(queue) > 0 {
			c := queue[0]
			queue = queue[1:]
			neighbours := getNeighbours(m, c)

			for _, n := range neighbours {
				x, y := n[0], n[1]
				if m[x][y] != 0 {
					m[x][y]++
					if m[x][y] > 9 {
						m[x][y] = 0
						flashes++
						queue = append(queue, newCell(x, y))
					}
				}
			}
		}
	}

	return flashes
}

func Part2(input []byte, steps int) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	var m matrix
	for scanner.Scan() {
		l := scanner.Text()
		n := strings.Split(l, "")

		var line []int
		for _, v := range n {
			num, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			line = append(line, num)
		}

		m = append(m, line)
	}

	for n := 0; n < steps; n++ {
		var flashes int

		for i, row := range m {
			for j := range row {
				m[i][j]++
			}
		}

		var queue []cell

		for i, row := range m {
			for j := range row {
				if row[j] > 9 {
					queue = append(queue, newCell(i, j))
					m[i][j] = 0
					flashes++
				}
			}
		}

		for len(queue) > 0 {
			c := queue[0]
			queue = queue[1:]
			neighbours := getNeighbours(m, c)

			for _, n := range neighbours {
				x, y := n[0], n[1]
				if m[x][y] != 0 {
					m[x][y]++
					if m[x][y] > 9 {
						m[x][y] = 0
						flashes++
						queue = append(queue, newCell(x, y))
					}
				}
			}
		}

		if flashes == len(m)*len(m[0]) {
			return n + 1
		}
	}
	return 0
}

func getNeighbours(matrix [][]int, centre [2]int) []cell {
	var neighbours []cell
	deltas := [][2]int{
		{-1, -1}, // top left
		{0, -1},  // top
		{1, -1},  // top right
		{-1, 0},  // left
		{1, 0},   // right
		{-1, 1},  // bottom left
		{0, 1},   // bottom
		{1, 1},   // bottom right
	}

	for _, d := range deltas {
		x, y := centre[0]+d[0], centre[1]+d[1]
		current := newCell(x, y)
		if x >= 0 && x < gridSize && y >= 0 && y < gridSize {
			neighbours = append(neighbours, current)
		}
	}

	return neighbours
}
