package day06

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/minaorangina/advent-of-code-2021/parse"
)

var (
	//go:embed testinput.txt
	testInput []byte
	//go:embed input.txt
	input []byte
)

func TestSolution(t *testing.T) {
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
		got := Solution(parse.ToIntSlice(testInput), tc.days)
		if got != tc.want {
			t.Errorf("got %d, want %d", got, tc.want)
		}
	}

	t.Run("real thing", func(t *testing.T) {
		fmt.Println(Solution(parse.ToIntSlice(input), 80))
		fmt.Println(Solution(parse.ToIntSlice(input), 256))
	})
}
