package day08

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
)

func Part1(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	var matrix [][]int

	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		matrix = append(matrix, []int{})

		for _, r := range []rune(line) {
			num, err := strconv.Atoi(string(r))
			if err != nil {
				log.Fatal(err)
			}

			matrix[row] = append(matrix[row], num)
		}

		row++
	}

	fmt.Println(matrix)

	perimeter := (len(matrix) * 2) + (len(matrix[0]) * 2) - 4

	count := perimeter

	for row := 1; row < len(matrix)-1; row++ {
		for col := 1; col < len(matrix[0])-1; col++ {
			coords := newCoord(row, col)
			if treeVisible(matrix, coords) {
				count++
			}
		}
	}

	return count
}

// visible right recursively (base: col == outer)
// visible left recursively (base: col == 0)
// visible up recursively (base: row == 0)
// visible down recursively (base: row == outer)

func treeVisible(matrix [][]int, c coord) bool {
	tree := c.value(matrix)

	fmt.Printf("start coord: (%v)\n . up (%v)\n . right (%v)\n . down (%v)\n . left (%v)\n",
		coord, row-1, col, row, col+1, row+1, col, row, col-1,
	)

	if matrix[row-1][col] < tree {
		return true
	}
	if matrix[row][col+1] < tree {
		return true
	}
	if matrix[row+1][col] < tree {
		return true
	}
	if matrix[row][col-1] < tree {
		return true
	}

	return false
}

type coord [2]int

func newCoord(x, y int) coord {
	return [2]int{x, y}
}
func (c coord) x() int {
	return c[0]
}
func (c coord) y() int {
	return c[1]
}
func (c coord) value(m [][]int) int {
	return m[c.x()][c.y()]
}
func (c coord) up(m [][]int) int {
	return m[c.x()-1][c.y()]
}
func (c coord) down(m [][]int) int {
	return m[c.x()+1][c.y()]
}
func (c coord) left(m [][]int) int {
	return m[c.x()][c.y()-1]
}
func (c coord) right(m [][]int) int {
	return m[c.x()][c.y()+1]
}
