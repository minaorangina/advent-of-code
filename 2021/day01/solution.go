package day01

func Part1(input []int) int {
	var count int

	for i, v := range input {
		if i == 0 {
			continue
		}
		if v > input[i-1] {
			count++
		}
	}

	return count
}

func Part2(input []int) int {
	var count int

	for i := range input {
		if i == 0 {
			continue
		}
		if i+2 >= len(input) {
			break
		}

		var sumWindow, sumPrevWindow int
		sumWindow += input[i] + input[i+1] + input[i+2]
		sumPrevWindow += input[i-1] + input[i] + input[i+1]

		if sumWindow > sumPrevWindow {
			count++
		}
	}

	return count
}
