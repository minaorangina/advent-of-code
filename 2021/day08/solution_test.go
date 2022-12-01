package day08

import (
	_ "embed"
	"fmt"
	"testing"
)

var (
	//go:embed testinput.txt
	testInput []byte
	//go:embed input.txt
	input []byte
)

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	if got != 26 {
		t.Fatalf("got %d, want 26", got)
	}

	t.Run("real thing", func(t *testing.T) {
		fmt.Println(Part1(input))
	})
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)
	if got != 61229 {
		t.Fatalf("got %d, want 61229", got)
	}
}
