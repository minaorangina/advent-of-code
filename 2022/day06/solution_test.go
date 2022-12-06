package day06

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed input.txt
var input []byte

func TestSolution(t *testing.T) {
	tt := []struct {
		input     []byte
		charCount int
		want      int
	}{
		{
			input:     []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"),
			want:      7,
			charCount: 4,
		},
		{
			input:     []byte("bvwbjplbgvbhsrlpgdmjqwftvncz"),
			want:      5,
			charCount: 4,
		},
		{
			input:     []byte("nppdvjthqldpwncqszvftbrmjlhg"),
			want:      6,
			charCount: 4,
		},
		{
			input:     []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
			want:      10,
			charCount: 4,
		},
		{
			input:     []byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
			want:      11,
			charCount: 4,
		},
		{
			input:     []byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"),
			want:      19,
			charCount: 14,
		},
		{
			input:     []byte("bvwbjplbgvbhsrlpgdmjqwftvncz"),
			want:      23,
			charCount: 14,
		},
		{
			input:     []byte("nppdvjthqldpwncqszvftbrmjlhg"),
			want:      23,
			charCount: 14,
		},
		{
			input:     []byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
			want:      29,
			charCount: 14,
		},
		{
			input:     []byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
			want:      26,
			charCount: 14,
		},
	}

	for _, tc := range tt {
		if res := Solution(tc.input, tc.charCount); res != tc.want {
			t.Errorf("want %d, got %d", tc.want, res)
		}
	}

	fmt.Println(Solution(input, 4))
	fmt.Println(Solution(input, 14))
}
