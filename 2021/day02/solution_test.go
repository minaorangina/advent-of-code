package day02

import (
	_ "embed"
	"fmt"
	"testing"
)

var (
	//go:embed input.txt
	input []byte
	//go:embed testinput.txt
	testInput []byte
)

func TestPart1(t *testing.T) {
	t.Run("test case", func(t *testing.T) {
		res, err := Part1(testInput)
		if err != nil {
			t.Fatal(err)
		}
		if res != 150 {
			t.Fatalf("wanted 150, got %d", res)
		}
	})

	t.Run("real thing", func(t *testing.T) {
		res, err := Part1(input)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("Day 2, part 1: %d", res)
	})
}
func TestPart2(t *testing.T) {
	t.Run("test case", func(t *testing.T) {
		res, err := Part2(testInput)
		if err != nil {
			t.Fatal(err)
		}
		if res != 900 {
			t.Fatalf("wanted 900, got %d", res)
		}
	})

	t.Run("real thing", func(t *testing.T) {
		res, err := Part2(input)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Printf("Day 2, part 2: %d", res)
	})
}
