package main

import (
	"fmt"
	"log"

	"advent-of-code-2024.com/internal/shared"
)

func main() {
	part1, err := wordSearchForXmas("testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Count XMAS: %v\n", part1)
	part2, err := wordSearchForMasXed("testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Count MAS X'd: %v\n", part2)
}

func wordSearchForXmas(filename string) (int, error) {
	nodes, err := wordSearchForString(filename, "XMAS" /*diagonalOnly=*/, false)
	if err != nil {
		return -1, err
	}
	return len(nodes), nil
}

func wordSearchForMasXed(filename string) (int, error) {
	nodes, err := wordSearchForString(filename, "MAS" /*diagonalOnly=*/, true)
	if err != nil {
		return -1, err
	}
	seenAs := map[coordinate]struct{}{}
	count := 0
	for _, seq := range nodes {
		// previous should always be the A
		if _, ok := seenAs[*seq.previous.coordinate]; ok {
			count++
		} else {
			seenAs[*seq.previous.coordinate] = struct{}{}
		}
	}
	return count, nil
}

func wordSearchForString(filename string, str string, diagonalOnly bool) ([]*node, error) {
	input, err := shared.FileToRuneMatrix(filename)
	if err != nil {
		return nil, err
	}
	queue := []*node{}
	for i := 0; i < len(input); i++ {
		line := input[i]
		for j := 0; j < len(line); j++ {
			queue = append(queue, &node{
				coordinate: &coordinate{
					i: i,
					j: j,
				},
				expectedChar: rune(str[0]), // depth is 0
			})
		}
	}
	var validSequences []*node
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		// don't bother validating coordinates until now
		if item.i < 0 || item.j < 0 {
			continue
		}
		if item.i >= len(input) || item.j >= len(input[item.i]) {
			continue
		}
		if input[item.i][item.j] != item.expectedChar {
			continue
		}
		depth := item.findDepth()
		if depth+1 >= len(str) {
			validSequences = append(validSequences, item)
			continue
		}
		nextExpectedChar := rune(str[depth+1])
		if depth == 0 {
			indices := []int{-1, 1}
			if !diagonalOnly {
				indices = append(indices, 0)
			}
			for _, i := range indices {
				for _, j := range indices {
					queue = append(queue, &node{
						coordinate: &coordinate{
							i: item.i + i,
							j: item.j + j,
						},
						expectedChar: nextExpectedChar,
						previous:     item,
					})
				}
			}
		} else {
			deltaI, deltaJ := item.difference(item.previous)
			queue = append(queue, &node{
				coordinate: &coordinate{
					i: item.i + deltaI,
					j: item.j + deltaJ,
				},
				expectedChar: nextExpectedChar,
				previous:     item,
			})
		}
	}
	return validSequences, nil
}

type coordinate struct {
	i, j int
}

type node struct {
	*coordinate
	expectedChar rune
	previous     *node
}

func (c *node) findDepth() int {
	if c.previous == nil {
		return 0
	}
	return 1 + c.previous.findDepth()
}

func (c1 *node) difference(c2 *node) (int, int) {
	return c1.i - c2.i, c1.j - c2.j
}
