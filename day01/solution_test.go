package day01

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/minaorangina/advent-of-code-2021/parse"
)

//go:embed input.txt
var input []byte

func TestPart1(t *testing.T) {
	testInput := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	got := Part1(testInput)

	if got != 7 {
		t.Error()
	}

	t.Run("real thing", func(t *testing.T) {
		input := parse.ToIntSlice(input)
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
		input := parse.ToIntSlice(input)
		fmt.Println(Part2(input))
	})
}
