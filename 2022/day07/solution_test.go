package day07

import (
	"bytes"
	_ "embed"
	"fmt"
	"testing"
)

//go:embed testinput.txt
var testinput []byte

//go:embed input.txt
var input []byte

func TestPart1(t *testing.T) {
	expected1, expected2 := 95437, 24933642

	testPart1, testPart2 := Solution(bytes.NewReader(testinput))

	if testPart1 != expected1 {
		t.Errorf("PART 1: got %d, want %d", testPart1, expected1)
	}

	if testPart2 != expected2 {
		t.Errorf("PART 2: got %d, want %d", testPart2, expected2)
	}

	part1, part2 := Solution(bytes.NewReader(input))
	fmt.Println("actual answers:", part1, part2)
}
