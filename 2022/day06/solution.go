package day06

func Solution(b []byte, charCount int) int {

	r := []rune(string(b))

	var answer int

	for i := charCount; i <= len(r); i++ {
		subslice := r[i-charCount : i]
		seen := map[rune]struct{}{}

		for j := 0; j < len(subslice); j++ {
			seen[subslice[j]] = struct{}{}
		}

		if len(seen) == charCount {
			answer = i
			break
		}
	}

	return answer
}
