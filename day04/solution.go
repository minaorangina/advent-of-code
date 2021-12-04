package day04

import (
	"log"
	"strconv"
	"strings"
)

func Part1(input []byte) int {
	game := bingo{}
	game.init(input)
	return game.play()
}

/*
need a way to determine a horizontal or vertical row of numbers

win state:
for a board, 5 numbers with the same horizontal or vertical coordinate

identifier per row/column, per board
- pull number, update row/column identifiers per board
- scan through and look for 5

sum of all unmarked numbers
- calculate at the beginning, decrement as we go

finding a number on the board



*/

type position struct {
	row, column int
}

type board struct {
	positions map[int]position
	unmarked  int
	called    map[string]int
}

type bingo struct {
	boards []board
	toCall []int
}

func (b *bingo) init(data []byte) {
	// parse
	b.parse(data)
	// setup positions for each board, calculate unmarked
}

func (b *bingo) play() int {
	// take next number
	// for each board, find position, update called and decrement unmarked
	// break on win condition

	// calculate result
	return 0
}

func (b *bingo) parse(input []byte) {
	d := strings.Split(string(input), "\n\n")

	for _, c := range d[0] {
		if string(c) == "," {
			continue
		}
		num, err := strconv.Atoi(string(c))
		if err != nil {
			log.Fatal(err)
		}
		b.toCall = append(b.toCall, num)
	}

	for boardNum, brd := range d[1:] {
		b.boards = append(b.boards, board{
			positions: map[int]position{},
			called:    map[string]int{},
		})

		rows := strings.Split(brd, "\n")

		for rowIdx, r := range rows {
			if r == "" {
				continue
			}
			asSlice := strings.Fields(r)

			var unmarked int
			for colIdx, c := range asSlice {
				n, err := strconv.Atoi(c)
				if err != nil {
					log.Fatal(err)
				}

				b.boards[boardNum].positions[n] = position{
					row: rowIdx, column: colIdx}

				unmarked += n
			}

			b.boards[boardNum].unmarked = unmarked
		}
	}
}
