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
	expected := 95437
	if res := Part1(bytes.NewReader(testinput)); res != expected {
		t.Errorf("got %d, want %d", res, expected)
	}

	answer := Part1(bytes.NewReader(input))
	fmt.Println("actual answer", answer)
}
