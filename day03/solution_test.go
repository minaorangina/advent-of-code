package day03

import (
	_ "embed"
	"testing"

	"github.com/minaorangina/advent-of-code-2021/parse"
)

//go:embed testinput.txt
var testInput []byte

func TestPart1(t *testing.T) {
	got := Part1(parse.ToIntSlice(testInput))

	if got != 198 {
		t.Fatalf("got %d, want 198", got)
	}
}
