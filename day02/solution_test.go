package day02

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	t.Run("test case", func(t *testing.T) {
		testInput := []byte(`forward 5
		down 5
		forward 8
		up 3
		down 8
		forward 2`)

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
		testInput := []byte(`forward 5
		down 5
		forward 8
		up 3
		down 8
		forward 2`)

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
