package day06

import (
	_ "embed"
	"testing"
)

var (
	//go:embed testinput.txt
	testInput []byte
	//go:embed input.txt
	input []byte
)

func TestPart1(t *testing.T) {
	tt := []struct {
		days, want int
	}{
		{
			days: 18,
			want: 26,
		},
		{
			days: 80,
			want: 5934,
		},
		{
			days: 256,
			want: 26984457539,
		}}

	for _, tc := range tt {
		got := Part1(testInput, tc.days)
		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}

	// t.Run("real thing", func(t *testing.T) {
	// 	fmt.Println(Part1(input, 80))
	// })
}
