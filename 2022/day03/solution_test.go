package day03

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
	if res := Part1(bytes.NewReader(testinput)); res != 157 {
		t.Errorf("want 157, got %d", res)
	}
	fmt.Println(Part1(bytes.NewReader(input)))
}

func TestPart2(t *testing.T) {
	if res := Part2(bytes.NewReader(testinput)); res != 70 {
		t.Errorf("want 70, got %d", res)
	}
	fmt.Println(Part2(bytes.NewReader(input)))
}
