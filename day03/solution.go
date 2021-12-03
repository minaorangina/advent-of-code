package day03

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func Part1(rows []string) int {
	numColumns := len(rows[0])
	oneTracker := make([]int, numColumns)

	for _, r := range rows {
		toLookAt := r
		for i, char := range toLookAt {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}

			oneTracker[i] += num
		}
	}

	var gammaRate, epsilonRate uint
	for i, v := range oneTracker {
		binaryIdx := numColumns - i

		if v > len(rows)/2 {
			gammaRate += uint(math.Pow(2, float64(binaryIdx)-1))
		} else {
			epsilonRate += uint(math.Pow(2, float64(binaryIdx)-1))
		}
	}

	return int(gammaRate * epsilonRate)
}

func Part2(rows []string) int {
	numColumns := len(rows[0])
	oneTracker, rowTracker := make([]int, numColumns), make([]int, numColumns)
	var oxygen, co2 uint

	// the number of rows changes as we move to the right
	// need to go column by column
	check := []int{}

	for col := 0; col < numColumns; col++ {
		var ones, zeros []string

		for r := 0; r < len(rows); r++ {
			char := rows[r][col]
			num, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}
			oneTracker[col] += num
			if num == 1 {
				ones = append(ones, rows[r])
			} else {
				zeros = append(zeros, rows[r])
			}
		}
		rowTracker[col] += len(rows)

		// figure out the most populous here
		midpoint := float64(len(rows)) / 2
		binaryIdx := numColumns - col

		fmt.Printf("mid: %2f, col: %d, num 1s: %d\n", midpoint, col, oneTracker[col])

		if float64(oneTracker[col]) >= midpoint {
			check = append(check, 1)
			oxygen += uint(math.Pow(2, float64(binaryIdx)-1))
			rows = ones
		} else if float64(rowTracker[col]-oneTracker[col]) >= midpoint {
			check = append(check, 0)
			co2 += uint(math.Pow(2, float64(binaryIdx)-1))
			rows = zeros
		}
	}

	fmt.Println("onet", oneTracker, "rowt", rowTracker)

	fmt.Printf("check: %v\no2: %d %12b, c02: %d %12b", check, oxygen, oxygen, co2, co2)

	return int(oxygen * co2)
}
