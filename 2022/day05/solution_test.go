package day05

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed testinput.txt
var testinput []byte

//go:embed input.txt
var input []byte

func TestPart1(t *testing.T) {
	if res := Part1(testinput); res != "CMZ" {
		t.Errorf("got %s, want CMZ", res)
	}
	fmt.Println(Part1(input))
}

// TIL
// - breakout of an inner loop with labels
// 	 https://stackoverflow.com/questions/51996175/how-to-break-out-of-nested-loops-in-go
// - create a slice with length
