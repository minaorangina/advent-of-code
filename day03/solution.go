package day03

import (
	"fmt"
	"math"
	"strconv"
)

func Part1(rows []int) int {
	numColumns := len(strconv.Itoa(rows[0]))
	fmt.Println("len", numColumns, rows[0])
	oneTracker := make([]int, numColumns)

	for _, r := range rows {
		toLookAt := r
		var index int
		for toLookAt != 0 {
			remainder := toLookAt % 10
			toLookAt = toLookAt / 10

			oneTracker[numColumns-index-1] += remainder

			index++
		}
	}

	var gammaRate float64
	for i, v := range oneTracker {
		if v > numColumns {
			gammaRate += math.Pow(2, float64(numColumns-i-1))
		}
		if oneTracker[numColumns-1] == 1 {
			gammaRate += 1
		}
	}

	fmt.Println("gamma rate", gammaRate)

	// inverse
	// multiply
	return 0
}

// 10101

// 2^0 = 0
// 0 +1 = 1

// 2^column*num
// // 2^0+1=1
// 2^1*0=0
// 2^2*1=4
// 2^3*0=0
// 2^4*1=16
