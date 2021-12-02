package day02

import (
	"testing"
)

func TestPart1(t *testing.T) {
	testInput := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	if Part1([]byte(testInput)) != 150 {
		t.Fatal()
	}
}
