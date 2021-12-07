package day04

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Solution(input []byte, isPart1 bool) int {
	game := bingo{}
	game.parse(input)
	return game.play(isPart1)
}

type position struct {
	row, column int
}

type board struct {
	positions           map[int]position
	unmarked            int
	winTracker          map[string]int
	winningBoardTracker int
}

func (b board) won() bool {
	for _, markedCount := range b.winTracker {
		if markedCount == 5 {
			return true
		}
	}
	return false
}

func (b board) wonPart2(numBoards int) bool {
	for _, markedCount := range b.winTracker {
		if markedCount == 5 {
			b.winningBoardTracker++
			if b.winningBoardTracker == numBoards {
				return true
			}
		}
	}
	return false
}

func (b board) score(lastDrawn int) int {
	return b.unmarked * lastDrawn
}

type bingo struct {
	boards []*board
	toCall []int
}

func (b *bingo) play(isPart1 bool) int {
	for _, drawn := range b.toCall {
		for _, brd := range b.boards {
			pos, ok := brd.positions[drawn]
			if !ok {
				continue
			}
			brd.unmarked -= drawn

			colKey := fmt.Sprintf("col-%d", pos.column)
			rowKey := fmt.Sprintf("row-%d", pos.row)
			brd.winTracker[colKey]++
			brd.winTracker[rowKey]++

			if isPart1 {
				if brd.won() {
					return brd.score(drawn)
				}
			} else if brd.wonPart2(len(b.boards)) {
				return brd.score(drawn)
			}

		}
	}

	return 0
}

func (b *bingo) parse(input []byte) {
	d := strings.Split(string(input), "\n\n")
	callingNums := strings.Split(d[0], ",")

	for _, char := range callingNums {
		num, err := strconv.Atoi(char)
		if err != nil {
			log.Fatal(err)
		}
		b.toCall = append(b.toCall, num)
	}

	for boardNum, brd := range d[1:] {
		b.boards = append(b.boards, &board{
			positions:  map[int]position{},
			winTracker: map[string]int{},
		})

		var unmarked int
		rows := strings.Split(brd, "\n")

		for rowIdx, r := range rows {
			if r == "" {
				continue
			}

			numbers := strings.Fields(r)

			for colIdx, c := range numbers {
				n, err := strconv.Atoi(c)
				if err != nil {
					log.Fatal(err)
				}

				b.boards[boardNum].positions[n] = position{
					row: rowIdx, column: colIdx}

				unmarked += n
			}
		}

		b.boards[boardNum].unmarked = unmarked
	}
}
