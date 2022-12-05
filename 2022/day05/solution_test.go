package day05

import (
	"bytes"
	_ "embed"
	"testing"
)

//go:embed testinput.txt
var testinput []byte

func TestPart1(t *testing.T) {
	if res := Part1(bytes.NewReader(testinput)); res != "CMZ" {
		t.Errorf("got %s, want CMZ", res)
	}
}
