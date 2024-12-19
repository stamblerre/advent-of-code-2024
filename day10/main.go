package main

import (
	"fmt"
	"log"

	"advent-of-code-2024.com/internal/shared"
)

func main() {
	t := &today{}
	_, _, err := shared.Run(t, "testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
}

type today struct {
}

func (t *today) ReadInput(filename string) (any, error) {
	return shared.ReadIntMatrix(filename)
}

func (t *today) Part1(input any) (int, error) {
	return implementation(input, 1)
}

func (t *today) Part2(input any) (int, error) {
	return implementation(input, 2)
}

type node struct {
	shared.Coordinate
	parent *node

	// always hold a reference to the trailhead for quick reference
	trailhead *node
}

type trail struct {
	trailhead shared.Coordinate
	trailend  shared.Coordinate
}

func implementation(input any, part int) (int, error) {
	matrix, ok := input.([][]int)
	if !ok {
		return -1, fmt.Errorf("unexpected input type %T", input)
	}
	// Add all potential trailheads.
	var queue []*node
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, &node{
					Coordinate: shared.Coordinate{
						I: i,
						J: j,
					},
				})
			}
		}
	}
	var trailEnds []*node
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		value := matrix[n.I][n.J]

		// end of the trail!
		if value == 9 {
			trailEnds = append(trailEnds, n)
		} else {
			// if no trailhead, then this is the trailhead
			trailhead := n.trailhead
			if trailhead == nil {
				trailhead = n
			}
			for _, delta := range shared.CardinalDirectionDelta() {
				next := n.Coordinate.Add(delta)

				if !shared.InBounds(matrix, *next) {
					continue
				}

				// need to increase by 1
				if matrix[next.I][next.J] == value+1 {
					queue = append(queue, &node{
						Coordinate: *next,
						parent:     n,
						trailhead:  trailhead,
					})
				}
			}
		}
	}

	switch part {
	case 1:
		scores := map[shared.Coordinate]int{}
		seen := map[trail]struct{}{}
		for _, end := range trailEnds {
			t := trail{
				trailhead: end.trailhead.Coordinate,
				trailend:  end.Coordinate,
			}
			if _, ok := seen[t]; ok {
				continue
			}
			seen[t] = struct{}{}

			// increment score by one for each distinct trail end
			scores[end.trailhead.Coordinate]++
		}

		result := 0
		for _, score := range scores {
			result += score
		}
		return result, nil
	case 2:
		ratings := map[shared.Coordinate]int{}
		for _, end := range trailEnds {
			// +1 for each distinct trail
			ratings[end.trailhead.Coordinate]++
		}
		result := 0
		for _, score := range ratings {
			result += score
		}
		return result, nil
	default:
		return -1, nil
	}
}
