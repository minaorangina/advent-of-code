package day10

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
	if got != 26397 {
		t.Fatalf("got %d, want 26397", got)
	}

	t.Run("real thing", func(t *testing.T) {
		fmt.Println(Part1(input))
	})
}

func TestPart2(t *testing.T) {
	got := Part2(testInput)
	if got != 288957 {
		t.Fatalf("got %d, want 288957", got)
	}

	t.Run("real thing", func(t *testing.T) {
		fmt.Println(Part2(input))
	})
}
