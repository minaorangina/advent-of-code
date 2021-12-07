package day06

func Solution(input []int, days int) int {
	states := make([]int, 9)
	for _, v := range input {
		states[v]++
	}

	for d := 0; d < days; d++ {
		newStates := make([]int, 9)
		for i, v := range states {
			if v == 0 {
				continue
			}
			if i == 0 {
				newStates[6] = v
				newStates[8] = v
				continue
			}
			newStates[i-1] += v
		}
		states = newStates
	}

	var count int
	for _, v := range states {
		count += v
	}

	return count
}
