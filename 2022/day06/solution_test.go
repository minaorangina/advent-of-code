package day06

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input []byte

func TestPart1(t *testing.T) {
	tt := []struct {
		input []byte
		want  int
	}{
		{
			input: []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"),
			want:  7,
		},
		{
			input: []byte("bvwbjplbgvbhsrlpgdmjqwftvncz"),
			want:  5,
		},
		{
			input: []byte("nppdvjthqldpwncqszvftbrmjlhg"),
			want:  6,
		},
		{
			input: []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
			want:  10,
		},
		{
			input: []byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
			want:  11,
		},
	}

	for _, tc := range tt {
		if res := Part1(tc.input); res != tc.want {
			t.Errorf("want %d, got %d", tc.want, res)
		}
	}

	fmt.Println(Part1(input))
}
