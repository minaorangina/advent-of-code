package day12

import (
	_ "embed"
	"testing"
)

var (
	input []byte
	//go:embed testinput.txt
	testinput []byte
)

func TestPart1(t *testing.T) {
	Part1(testinput)
}
