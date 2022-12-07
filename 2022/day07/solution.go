package day07

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"strings"
)

type node struct {
	name     string
	entity   string // or enum?
	parent   *node
	children map[string]*node
	size     int // recursively answer this
}

type command []string

var errNotACommand = errors.New("not a command")

var (
	totalDiskSpace float64 = 70000000
	spaceRequired  float64 = 30000000
)

func Solution(r io.Reader) (int, int) {
	scanner := bufio.NewScanner(r)

	var root, currentNode *node

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(">>>", line)

		parts := strings.Split(line, " ")

		command, err := getCommand(parts)
		if errors.Is(err, errNotACommand) {
			// this must be a file/directory listing
			// we're partway through an ls

			// we have a child of the current node
			child := newNode(currentNode, parts)

			currentNode.children[child.name] = child
		}

		if command == "cd" {
			path, err := getDest(parts)
			if err != nil {
				log.Fatal(err)
			}

			switch path {
			case "/":
				r := newRootNode()
				currentNode = r
				root = r

			case "..":
				currentNode = currentNode.parent

			default: // some dir
				currentNode = currentNode.children[path]
			}
		}

		// fmt.Println("TREE")
		// printTree(root, 0)
		// fmt.Println("------")
	}

	_, ans := getSize(root)

	smallest := bestNodeToDelete(root, float64(root.size), 0)

	return ans, int(smallest)
}

func bestNodeToDelete(n *node, rootSize, num float64) float64 {
	// we've gone too far
	if n.entity == "file" {
		return math.Inf(1)
	}

	dirSizes := getSizes(n)
	bestSoFar := math.Inf(1)

	for _, s := range dirSizes {
		spaceRemaining := totalDiskSpace - rootSize + s

		if spaceRemaining >= spaceRequired && s < bestSoFar {
			bestSoFar = s
		}
	}

	return bestSoFar
}

func getSizes(n *node) []float64 {
	if n.entity == "file" {
		return nil
	}
	sizes := []float64{}

	sizes = append(sizes, float64(n.size))
	fmt.Println("adding", n.name, n.size)

	for _, c := range n.children {
		s := getSizes(c)
		sizes = append(sizes, s...)
	}

	return sizes
}

func getSize(n *node) (nodeSize int, sumSub100k int) {
	if n.entity == "file" {
		return n.size, 0
	}

	dirSize, collection := 0, 0
	for _, c := range n.children {
		s, c := getSize(c)
		dirSize += s
		collection += c
	}

	if dirSize <= 100000 {
		collection += dirSize
	}

	n.size = dirSize

	return dirSize, collection
}

func getDest(parts []string) (string, error) {
	if parts[1] == "cd" {
		return parts[2], nil
	}
	return "", fmt.Errorf(`not a cd command, got "%s"`, strings.Join(parts, " "))
}

func getCommand(parts []string) (string, error) {
	if parts[0] == "$" {
		return parts[1], nil
	}
	return "", fmt.Errorf(`%w: "%s"`, errNotACommand, strings.Join(parts, " "))
}

func newRootNode() *node {
	return &node{
		name:     "/",
		entity:   "dir",
		parent:   nil,
		children: map[string]*node{},
	}
}

func newNode(parent *node, parts []string) *node {
	entity := "file"
	if parts[0] == "dir" {
		entity = "dir"
	}

	n := &node{
		name:     parts[1],
		entity:   entity,
		parent:   parent,
		children: map[string]*node{},
	}

	if n.entity == "file" {
		size, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		n.size = size
	}

	return n
}

func printTree(n *node, level int) {
	msg := fmt.Sprintf("%s- %s", strings.Repeat("  ", level), n.name)

	msg += fmt.Sprintf(" (%s, size=%d)", n.entity, n.size)

	if n.entity == "dir" {
		msg += " (dir)"
		msg += fmt.Sprintf(" [numChildren: %d]", len(n.children))
	}

	fmt.Println(msg)

	if n.entity == "dir" {
		for _, child := range n.children {
			printTree(child, level+1)
		}
	}
}
