package day08

import (
	"bytes"
	"fmt"
	"testing"

	_ "embed"
)

var (
	//go:embed testinput.txt
	testinput []byte
	//go:embed input.txt
	input []byte
)

func TestPart1(t *testing.T) {
	expected := 21

	if ans := Part1(bytes.NewReader(testinput)); ans != expected {
		t.Errorf("got %d, want %d", ans, expected)
	}

	fmt.Println(Part1(bytes.NewReader(input)))
}
func TestPart2(t *testing.T) {
	expected := 8

	if ans := Part2(bytes.NewReader(testinput)); ans != expected {
		t.Errorf("got %d, want %d", ans, expected)
	}

	fmt.Println(Part2(bytes.NewReader(input)))
}
