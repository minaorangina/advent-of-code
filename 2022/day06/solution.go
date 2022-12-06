package day06

func Part1(b []byte) int {

	r := []rune(string(b))

	var answer int

	for i := 4; i <= len(r); i++ {
		subslice := r[i-4 : i]
		seen := map[rune]struct{}{}

		for j := 0; j < len(subslice); j++ {
			seen[subslice[j]] = struct{}{}
		}

		if len(seen) == 4 {
			answer = i
			break
		}
	}

	return answer
}
