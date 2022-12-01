package day04

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
	got := Solution(testInput, true)

	if got != 4512 {
		t.Errorf("got %d, want 4512", got)
	}

	t.Run("real thing", func(t *testing.T) {
		fmt.Println(Solution(input, true))
	})
}
func TestPart2(t *testing.T) {
	got := Solution(testInput, false)

	if got != 4512 {
		t.Errorf("got %d, want 4512", got)
	}

	t.Run("real thing", func(t *testing.T) {
		fmt.Println(Solution(input, false))
	})
}
