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
	if res := Solution(testinput, fromBottom); res != "CMZ" {
		t.Errorf("got %s, want CMZ", res)
	}
	if res := Solution(testinput, fromTop); res != "MCD" {
		t.Errorf("got %s, want MCD", res)
	}

	fmt.Println(Solution(input, fromBottom))
	fmt.Println(Solution(input, fromTop))
}
