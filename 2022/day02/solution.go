package day02

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type outcome int

var (
	lose outcome = 0
	draw outcome = 3
	win  outcome = 6
)

var shapeScores = map[string]int{
	"A": 1, // rock
	"X": 1, // rock
	"B": 2, // paper
	"Y": 2, // paper
	"C": 3, // scissors
	"Z": 3, // scissors
}

var defeatTable = map[string]string{
	"A": "Z",
	"B": "X",
	"C": "Y",
}

func Part1(r io.Reader) int {
	var player2Score int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		raw := scanner.Text()
		line := strings.Split(raw, " ")

		shape2Score := shapeScores[line[1]]

		player2Score += shape2Score + int(getP2Score(line[0], line[1]))
	}
	return player2Score
}

func getP2Score(p1Choice, p2Choice string) outcome {
	if shapeScores[p1Choice] == shapeScores[p2Choice] {
		fmt.Println(p1Choice, p2Choice, "draw!")
		return draw
	}
	if defeatTable[p1Choice] == p2Choice {
		fmt.Println(p1Choice, p2Choice, "lose!")
		return lose
	}
	fmt.Println(p1Choice, p2Choice, "win!")
	return win
}

var offsets = map[string]int{
	"X": -1, // lose, look behind
	"Y": 0,  // draw
	"Z": 1,  // win, look ahead
}

var rules = map[int]string{
	-1: "lose",
	0:  "draw",
	1:  "win",
}

var p1Choices = []string{"A", "B", "C"}
var p2Choices = []string{"X", "Y", "Z"}

func Part2(r io.Reader) int {
	var player2Score int

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		raw := scanner.Text()
		line := strings.Split(raw, " ")

		p1Choice := line[0]

		offset := offsets[line[1]]

		p1Idx := indexOf(p1Choices, p1Choice)
		p2Idx := ((p1Idx + offset) + len(p1Choices)) % len(p1Choices)

		fmt.Printf("should %s: offset=%d p1idx %d, p2idx %d",
			rules[offset], offset, p1Idx, p2Idx)

		p2Choice := p2Choices[p2Idx]

		fmt.Printf(" -- (%s vs %s) -> ", p1Choice, p2Choice)

		player2Score += shapeScores[p2Choice] + int(getP2Score(p1Choice, p2Choice))
	}
	return player2Score
}

func indexOf(s []string, choice string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == choice {
			return i
		}
	}
	return -1
}
