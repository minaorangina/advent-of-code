package day10

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

func Part1(r io.Reader) int {
	scanner := bufio.NewScanner(r)

	cycleToCheck, interval := 20, 40
	x, signalStrength := 1, 0

	cycle := 1

	for scanner.Scan() {
		// if cycle == 180 || cycle == 140 || cycle == 100 {
		// 	fmt.Println("here")
		// }

		// check numbers before continuing
		if cycle == cycleToCheck {
			ss := (cycle * x)
			signalStrength += ss
			cycleToCheck += interval
		}

		instr := scanner.Text()
		if instr == "noop" {

			// check numbers before continuing
			if cycle == cycleToCheck {
				ss := (cycle * x)
				signalStrength += ss
				cycleToCheck += interval
			}

			cycle++

			continue
		}

		addX := strings.Split(instr, " ")
		toAdd, err := strconv.Atoi(addX[1])
		if err != nil {
			log.Fatal(err)
		}

		cycle++

		// if cycle == 180 || cycle == 140 || cycle == 100 {
		// 	fmt.Println("here")
		// }

		// check numbers before continuing
		if cycle == cycleToCheck {
			ss := (cycle * x)
			signalStrength += ss
			cycleToCheck += interval
		}

		x += toAdd

		cycle++
	}

	return signalStrength
}
