package day03

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/minaorangina/advent-of-code-2021/parse"
)

var (
	//go:embed testinput.txt
	testInput []byte
	//go:embed input.txt
	input []byte
)

func TestPart1(t *testing.T) {
	got := Part1(parse.ToStringSlice(testInput))

	if got != 198 {
		t.Fatalf("got %d, want 198", got)
	}

	t.Run("real thing", func(t *testing.T) {
		fmt.Println(Part1(parse.ToStringSlice(input)))
	})
}
func TestPart2(t *testing.T) {
	got := Part2(parse.ToStringSlice(testInput))

	if got != 230 {
		t.Fatalf("got %d, want 230", got)
	}
	t.Run("real thing", func(t *testing.T) {
		fmt.Println(Part1(parse.ToStringSlice(input)))
	})
}
