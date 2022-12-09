package day09

import (
	"io"
	"math"
	"strings"
)

func Part1(r io.Reader) {}

func makeMatrix(steps []step) matrix {
	minX, maxX, cursorX := 0, 0, 0
	minY, maxY, cursorY := 0, 0, 0

	for _, s := range steps {
		if s.dir == "R" {
			cursorX = s.dist

			if cursorX > maxX {
				maxX = cursorX
			}
		}

		if s.dir == "L" {
			cursorX -= s.dist

			if cursorX < minX {
				minX = cursorX
			}
		}

		if s.dir == "U" {
			cursorY += s.dist

			if cursorY > maxY {
				maxY = cursorY
			}
		}

		if s.dir == "D" {
			cursorY -= s.dist

			if cursorY < minY {
				minY = cursorY
			}
		}
	}

	lengthX := maxX + int(math.Abs(float64(minX))) + 1
	// y axis will be inverted
	lengthY := maxY + int(math.Abs(float64(minY))) + 1

	startX := lengthX - maxX - 1
	startY := lengthY - maxY - 1

	m := newMatrix(lengthX, lengthY)
	m.setStart(newCoord(startX, startY))

	return m
}

type step struct {
	dir  string
	dist int
}

type matrix struct {
	grid  [][]string
	start coord
}

func newMatrix(lengthX, lengthY int) matrix {
	var m [][]string

	for y := 0; y < lengthY; y++ {
		m = append(m, []string{})

		for x := 0; x < lengthX; x++ {
			m[y] = append(m[y], ".")
		}
	}

	return matrix{
		grid: m,
	}
}

func (m matrix) setStart(c coord) {
	m.start = c
	m.grid[c.y()][c.x()] = "s"
}

func (m matrix) String() string {
	var s strings.Builder

	for _, row := range m.grid {
		for _, v := range row {
			s.WriteString(v)
		}

		s.WriteString("\n")
	}

	return s.String()
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
