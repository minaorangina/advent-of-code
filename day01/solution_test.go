package day01

import (
	"fmt"
	"testing"

	"github.com/minaorangina/advent-of-code-2021/data"
)

func TestPart1(t *testing.T) {
	testInput := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	got := Part1(testInput)

	if got != 7 {
		t.Error()
	}

	t.Run("real thing", func(t *testing.T) {
		input := data.ToIntSlice(data.Parse())
		fmt.Println(Part1(input))
	})
}

func TestPart2(t *testing.T) {
	testInput := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	got := Part2(testInput)

	if got != 5 {
		t.Error()
	}

	t.Run("real thing", func(t *testing.T) {
		input := data.ToIntSlice(data.Parse())
		fmt.Println(Part2(input))
	})
}
