package day02

import (
	"bytes"
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
	if res := Part1(bytes.NewReader(testInput)); res != 15 {
		t.Errorf("got %d, want 15", res)
	}

	fmt.Println(Part1(bytes.NewReader(input)))
}

func TestPart2(t *testing.T) {
	if res := Part2(bytes.NewReader(testInput)); res != 12 {
		t.Errorf("got %d, want 12", res)
	}

	fmt.Println(Part2(bytes.NewReader(input)))
}
