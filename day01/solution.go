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
	return 0
}
