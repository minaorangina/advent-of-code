package day11

import (
	_ "embed"
	"fmt"
	"math"
	"testing"
)

var (
	//go:embed testinput.txt
	testInput []byte
	//go:embed input.txt
	input []byte
)

func TestSolution(t *testing.T) {
	tt := []struct {
		steps, want int
	}{
		{steps: 10, want: 204},
		{steps: 100, want: 1656},
	}

	for _, tc := range tt {
		got := Part1(testInput, tc.steps)
		if got != tc.want {
			t.Fatalf("got %d, want %d", got, tc.want)
		}
	}

	t.Run("real thing", func(t *testing.T) {
		fmt.Println(Part1(input, 100))
	})

	t.Run("part 2", func(t *testing.T) {
		fmt.Println(Part2(input, math.MaxInt))
	})
}
