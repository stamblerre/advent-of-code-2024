package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	result, err := wordSearch("testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Count: %v\n", result)
}

var depthToCharMap = map[int]rune{
	0: 'X',
	1: 'M',
	2: 'A',
	3: 'S',
}

func wordSearch(filename string) (int, error) {
	text, err := os.ReadFile(filename)
	if err != nil {
		return -1, err
	}
	// convert to 2D array
	var input [][]rune
	for i, line := range strings.Split(string(text), "\n") {
		input = append(input, []rune{})
		for _, c := range line {
			input[i] = append(input[i], c)
		}
	}
	queue := []*coordinate{}
	for i := 0; i < len(input); i++ {
		line := input[i]
		for j := 0; j < len(line); j++ {
			queue = append(queue, &coordinate{
				i:            i,
				j:            j,
				expectedChar: depthToCharMap[0], // depth is 0
			})
		}
	}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		// don't bother validating coordinates until now
		if item.i >= len(input) || item.j >= len(input[item.i]) {
			continue
		}
		if input[item.i][item.j] != item.expectedChar {
			continue
		}
		depth := item.findDepth()
		if depth == 0 {
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {

				}
			}
		}
	}
	return -1, nil
}

type coordinate struct {
	i, j         int
	expectedChar rune
	previous     *coordinate
}

func (c *coordinate) findDepth() int {
	if c.previous == nil {
		return 0
	}
	return 1 + c.previous.findDepth()
}
