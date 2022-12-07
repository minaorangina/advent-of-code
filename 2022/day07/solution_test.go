package day07

import (
	"bytes"
	_ "embed"
	"testing"
)

//go:embed testinput.txt
var testinput []byte

func TestPart1(t *testing.T) {
	expected := 95437
	if res := Part1(bytes.NewReader(testinput)); res != expected {
		t.Errorf("got %d, want %d", res, expected)
	}
}
