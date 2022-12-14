package day01

import (
	"bytes"
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
	if res := Part1(bytes.NewReader(testInput)); res != 24000 {
		t.Errorf("got %d, want 24000", res)
	}

	fmt.Println(Part1(bytes.NewReader(input)))
}

func TestPart2(t *testing.T) {
	if res := Part2(bytes.NewReader(testInput)); res != 45000 {
		t.Errorf("got %d, want 45000", res)
	}

	fmt.Println(Part2(bytes.NewReader(input)))
}
