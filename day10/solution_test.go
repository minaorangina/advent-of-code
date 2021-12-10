package day10

import (
	_ "embed"
	"testing"
)

var (
	//go:embed testinput.txt
	testInput []byte
	input     []byte
)

func TestPart1(t *testing.T) {
	got := Part1(testInput)
	if got != 26397 {
		t.Fatalf("got %d, want 26397", got)
	}
}
