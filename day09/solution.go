package day09

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"sync"
)

func Part1(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	var outer [][]int
	for scanner.Scan() {
		line := scanner.Text()

		var inner []int
		for _, c := range line {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			inner = append(inner, num)
		}
		outer = append(outer, inner)
	}

	var riskLevel int
	for i, inner := range outer {
		for j, v := range inner {
			// left inner[j-1]
			leftIdx := j - 1
			if leftIdx >= 0 {
				if inner[leftIdx] <= v {
					continue
				}
			}
			// right inner[j+1]
			rightIdx := j + 1
			if rightIdx < len(inner) {
				if inner[rightIdx] <= v {
					continue
				}
			}
			// up outer[i-1][j]
			topIdx := i - 1
			if topIdx >= 0 {
				if outer[topIdx][j] <= v {
					continue
				}
			}
			// down outer[i+1][j]
			bottomIdx := i + 1
			if bottomIdx < len(outer) {
				if outer[bottomIdx][j] <= v {
					continue
				}
			}
			riskLevel += v + 1
		}
	}
	return riskLevel
}
func Part2(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	var outer [][]int
	for scanner.Scan() {
		line := scanner.Text()

		var inner []int
		for _, c := range line {
			num, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			inner = append(inner, num)
		}
		outer = append(outer, inner)
	}

	total := 1
	for i, inner := range outer {
		for j, v := range inner {
			// left inner[j-1]
			leftIdx := j - 1
			if leftIdx >= 0 {
				if inner[leftIdx] <= v {
					continue
				}
			}
			// right inner[j+1]
			rightIdx := j + 1
			if rightIdx < len(inner) {
				if inner[rightIdx] <= v {
					continue
				}
			}
			// up outer[i-1][j]
			topIdx := i - 1
			if topIdx >= 0 {
				if outer[topIdx][j] <= v {
					continue
				}
			}
			// down outer[i+1][j]
			bottomIdx := i + 1
			if bottomIdx < len(outer) {
				if outer[bottomIdx][j] <= v {
					continue
				}
			}

			// low point discovered
			c := cache{
				seen: map[string]struct{}{},
				mu:   &sync.Mutex{},
			}
			basinSize := find(outer, c, [2]int{i, j})
			if basinSize > 0 {
				total *= basinSize
			}
		}
	}
	return total
}

func find(matrix [][]int, c cache, start coord) int {
	outerMax, innerMax := len(matrix), len(matrix[0])
	if !start.valid(outerMax, innerMax) {
		return 0
	}
	if c.alreadySeen(start) {
		return 0
	}
	if start.value(matrix) == 9 {
		return 0
	}

	var count int
	c.mu.Lock()
	c.seen[start.String()] = struct{}{}
	count++
	c.mu.Unlock()

	count += find(matrix, c, start.up())
	count += find(matrix, c, start.down())
	count += find(matrix, c, start.left())
	count += find(matrix, c, start.right())
	fmt.Printf("count for %s: %d\n", start.String(), start.value(matrix))
	return count
}

type cache struct {
	seen map[string]struct{}
	mu   *sync.Mutex
}

func (c cache) alreadySeen(cd coord) bool {
	_, ok := c.seen[cd.String()]
	return ok
}

type coord [2]int

func (c coord) value(matrix [][]int) int {
	x, y := c[0], c[1]
	return matrix[x][y]
}
func (c coord) valid(outerMax, innerMax int) bool {
	return c[0] >= 0 && c[0] < outerMax &&
		c[1] >= 0 && c[1] < innerMax
}
func (c coord) String() string {
	return fmt.Sprintf("%d,%d", c[0], c[1])
}
func (c coord) up() coord {
	return [2]int{c[0] - 1, c[1]}
}
func (c coord) down() coord {
	return [2]int{c[0] + 1, c[1]}
}
func (c coord) left() coord {
	return [2]int{c[0], c[1] - 1}
}
func (c coord) right() coord {
	return [2]int{c[0], c[1] + 1}
}
