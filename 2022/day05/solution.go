package day05

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
)

func Part1(r io.Reader) string {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()

		res, err := regexp.MatchString("(\\d+\\s+)+", line)
		if err != nil {
			log.Fatal(err)
		}

		if res {
			// clean up
			fmt.Println([]rune(line))
			break
		}

		asRunes := []rune(line)

		for i := 1; i < len(asRunes); i = i + 4 {
			fmt.Println(string(asRunes[i]))
		}
	}

	return ""
}
