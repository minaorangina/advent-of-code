package main

import (
	"bytes"
	"testing"
)

type testCase struct {
	sample []byte
	want   int
}

func TestThreePart1(t *testing.T) {
	cases := []testCase{
		{
			sample: []byte(`..##.........##.........##.........##.........##.........##.......
#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....
.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........#.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...##....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#`),
			want: 7,
		},
		{
			sample: []byte(`.........###......#...#.......#
.#.#...........#..#..#.........
#.......#.................#....`),
			want: 1,
		},
	}

	for _, c := range cases {
		assertEqual(t,
			doThreePart1(bytes.NewReader(c.sample)),
			c.want)
	}
}
