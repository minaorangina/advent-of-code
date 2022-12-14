package day08

import (
	"bufio"
	"io"
	"log"
	"strconv"
)

func Part1(r io.Reader) int {
	matrix := getMatrix(r)

	perimeter := (len(matrix) * 2) + (len(matrix[0]) * 2) - 4
	count := perimeter

	for row := 1; row < len(matrix)-1; row++ {
		for col := 1; col < len(matrix[0])-1; col++ {
			coords := newCoord(row, col)

			if visibleAlongRow(matrix, coords, row) ||
				visibleDownColumn(matrix, coords, col) {
				count++
			}
		}
	}

	return count
}

func Part2(r io.Reader) int {
	matrix := getMatrix(r)

	bestScore := 0

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			coords := newCoord(row, col)

			score := latScenicScore(matrix, coords, row) *
				lngScenicScore(matrix, coords, col)

			if score > bestScore {
				bestScore = score
			}
		}
	}

	return bestScore
}

func getMatrix(r io.Reader) [][]int {
	scanner := bufio.NewScanner(r)

	var matrix [][]int

	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		matrix = append(matrix, []int{})

		for _, char := range string(line) {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}

			matrix[row] = append(matrix[row], num)
		}

		row++
	}

	return matrix
}

func latScenicScore(m [][]int, tree coord, rowIdx int) int {
	treeVal := tree.value(m)

	treeRow := m[rowIdx]

	left, right := treeRow[0:tree.y()], treeRow[tree.y()+1:]

	leftViewingDist := 0

	for i := range left {
		leftViewingDist++

		idx := tree.y() - 1 - i
		t := left[idx]

		if t >= treeVal {
			// can't see beyond this tree
			break
		}
	}

	rightViewingDist := 0

	for _, t := range right {
		rightViewingDist++

		if t >= treeVal {
			// can't see beyond this tree
			break
		}
	}

	return leftViewingDist * rightViewingDist
}

func lngScenicScore(m [][]int, tree coord, colIdx int) int {
	treeVal := tree.value(m)

	treeCol := []int{}

	for _, row := range m {
		treeCol = append(treeCol, row[colIdx])
	}

	up, down := treeCol[0:tree.x()], treeCol[tree.x()+1:]

	upperViewingDist := 0

	for i := range up {
		upperViewingDist++

		idx := tree.x() - 1 - i
		t := up[idx]

		if t >= treeVal {
			// can't see beyond this tree
			break
		}
	}

	lowerViewingDst := 0

	for _, t := range down {
		lowerViewingDst++

		if t >= treeVal {
			// can't see beyond this tree
			break
		}
	}

	return upperViewingDist * lowerViewingDst
}

func visibleAlongRow(m [][]int, tree coord, rowIdx int) bool {
	treeVal := tree.value(m)

	treeRow := m[rowIdx]

	left, right := treeRow[0:tree.y()], treeRow[tree.y()+1:]
	// fmt.Println("left", left, "right", right)

	var leftObscured bool
	for _, t := range left {
		if t >= treeVal {
			// view obscured to left
			leftObscured = true
			break
		}
	}

	// if left obscured, we should check the right
	// if it wasn't that's good enough
	if !leftObscured {
		return true
	}

	for _, t := range right {
		if t >= treeVal {
			// view obscured to right
			return false
		}
	}

	return true
}

func visibleDownColumn(m [][]int, tree coord, colIdx int) bool {
	treeVal := tree.value(m)

	treeCol := []int{}

	for _, row := range m {
		treeCol = append(treeCol, row[colIdx])
	}

	up, down := treeCol[0:tree.x()], treeCol[tree.x()+1:]

	// fmt.Println("up", up, "down", down)

	var topObscured bool
	for _, t := range up {
		if t >= treeVal {
			// view obscured above
			topObscured = true
			break
		}
	}

	// if top obscured, check bottom
	if !topObscured {
		return true
	}

	for _, t := range down {
		if t >= treeVal {
			// view obscured to below
			return false
		}
	}

	return true
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
