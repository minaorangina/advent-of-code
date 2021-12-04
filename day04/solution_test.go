package day04

import (
	_ "embed"
	"testing"
)

var (
	//go:embed testinput.txt
	testInput []byte
)

func TestPart1(t *testing.T) {
	got := Part1(testInput)

	if got != 4512 {
		t.Errorf("got %d, want 4512", got)
	}
}
