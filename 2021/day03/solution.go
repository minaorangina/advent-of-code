package day03

import (
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
	return 0
}
