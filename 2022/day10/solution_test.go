package day10

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
	want := 13140
	if got := Part1(bytes.NewReader(testinput)); got != want {
		t.Errorf("got %d, want %d", got, want)
	}

	fmt.Println(Part1(bytes.NewReader(input)))
}
