package day09

import (
	"bufio"
	"bytes"
	"log"
	"strconv"
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
