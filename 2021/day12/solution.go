package day12

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func Part1(input []byte) int {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	edgeList := map[string][]string{}

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")

		if _, ok := edgeList[line[0]]; ok {
			edgeList[line[0]] = append(edgeList[line[0]], line[1])
			continue
		}
		edgeList[line[0]] = []string{line[1]}
	}

	fmt.Println(edgeList)

	frontier := []path{{nodesVisited: []string{"start"}}}

	for len(frontier) > 0 {
		node := frontier[0]
		frontier = frontier[1:]
		children := node.getChildren(edgeList)
		frontier = append(frontier, children...)

	}
}

type path struct {
	nodesVisited []string
}

func (p path) getChildren(edgeList map[string][]string) []path {

	return nil
}

/*
- find all paths
- reject paths with more than one small cave
- DFS but backtrack when a second small cave is found


setup
- a map from node to edges
- a map of visited nodes
- some way to build the path as I go (and discard if necessary.) maybe a string.
*/
