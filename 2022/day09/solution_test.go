package day09

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed testinput.txt
var testinput []byte

func TestPart1(t *testing.T) {
	steps := []step{
		{"R", 2},
		{"L", 1},
		{"L", 2},
		{"U", 2},
		{"D", 1},
		{"D", 2},
	}

	m := makeMatrix(steps)
	fmt.Println(m.String())
}
