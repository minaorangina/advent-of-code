package day04

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
	if res := Part1(bytes.NewReader(testinput)); res != 2 {
		t.Errorf("want 2, got %d", res)
	}
	fmt.Println(Part1(bytes.NewReader(input)))
}
