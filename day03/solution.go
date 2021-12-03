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

const (
	least = iota
	most
)

func Part2(rows []string) int {
	numColumns := len(rows[0])

	var rowsA, rowsB []string

	rowsA = whittleDown(rows, 0, most)
	rowsB = whittleDown(rows, 0, least)
	fmt.Println("????", rowsA, rowsB)

	var oxygenRating, co2Rating uint16
	for i, v := range rowsA[0] {
		binaryIdx := numColumns - i

		numA, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("huh?", string(rowsB[0][1]))
		numB, err := strconv.Atoi(string(rowsB[0][i]))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("%%%%%", numA, numB)
		oxygenRating += uint16(math.Pow(2, float64(binaryIdx)-1) * float64(numA))
		co2Rating += uint16(math.Pow(2, float64(binaryIdx)-1) * float64(numB))
	}

	fmt.Println("oxy", oxygenRating, "co2", co2Rating)
	return int(oxygenRating * co2Rating)
}

func getMostCommon(oneCount, rowLength float64) byte {
	if oneCount >= rowLength/2 {
		return '1'
	}
	return '0'
}

func whittleDown(rows []string, idx int, searchFor int) []string {
	if len(rows) == 1 {
		return rows
	}
	oneTracker := 0
	// fmt.Printf("\nlen rows: %d\n", len(rows))

	for _, toLookAt := range rows {
		num, err := strconv.Atoi(string(toLookAt[idx]))
		if err != nil {
			log.Fatal(err)
		}

		oneTracker += num

	}
	// fmt.Printf("col %d we got %d 1s\n", idx, oneTracker)

	var winners []string
	mostCommon := getMostCommon(float64(oneTracker), float64((len(rows))))

	if searchFor == most {
		for _, r := range rows {
			if r[idx] == mostCommon {
				winners = append(winners, r)
			}
		}
	}

	if searchFor == least {
		for _, r := range rows {
			if r[idx] != mostCommon {
				winners = append(winners, r)
			}
		}
	}

	return whittleDown(winners, idx+1, searchFor)
}
