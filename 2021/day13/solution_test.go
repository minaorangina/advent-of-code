package day13

import (
	_ "embed"
	"fmt"
	"testing"
)

var (
	//go:embed input.txt
	input []byte
	//go:embed testinput.txt
	testInput []byte
)

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	if got != 17 {
		t.Errorf("got %d, want 17", got)
	}

	t.Run("real thing", func(t *testing.T) {
		fmt.Println(Part1(input))
	})
}
