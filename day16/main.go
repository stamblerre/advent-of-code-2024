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
	return shared.ReadRuneMatrix(filename)
}

func (t *today) Part1(input any) (int, error) {
	return implementation(input, 1)
}

func (t *today) Part2(input any) (int, error) {
	return implementation(input, 2)
}

type node struct {
	shared.Coordinate
	direction shared.Direction
	parent    *node

	// keep track of all the seen coordinates along the path
	seen map[shared.Coordinate]struct{}

	moveType // maybe...
}

func (n *node) String() string {
	return fmt.Sprintf("%v - %s", n.Coordinate, n.direction)
}

type moveType int

const (
	unknown moveType = iota

	step
	rotate
)

func implementation(input any, part int) (int, error) {
	matrix, ok := input.([][]rune)
	if !ok {
		return -1, fmt.Errorf("unexpected type %T for input", input)
	}
	start := shared.FindRune(matrix, 'S')
	end := shared.FindRune(matrix, 'E')

	stack := []*node{{
		Coordinate: *start,
		direction:  shared.Right,
	}}
	var paths []*node // pointer to the end
	for len(stack) > 0 {
		last := len(stack) - 1
		n := stack[last]
		stack = stack[:last]

		// reached the end!
		if n.Coordinate.Equals(end) {
			paths = append(paths, n)
			continue
		}

		// options:
		//
		// (1) go forward, if possible
		if next := n.AddDirection(n.direction); shared.InBounds(matrix, next) && !isWall(matrix, next) && !n.previouslySeen(next) {
			stack = append(stack, &node{
				direction:  n.direction,
				parent:     n,
				moveType:   step,
				Coordinate: n.AddDirection(n.direction),
				seen:       n.copySeenAndAdd(next),
			})
		}

		// (2) turn clockwise or counterclockwise
		// but don't turn if you've turned before...don't want to spin in circles
		if n.moveType != rotate {
			// clockwise
			stack = append(stack, &node{
				direction:  n.direction.Clockwise(),
				parent:     n,
				moveType:   rotate,
				Coordinate: n.Coordinate,
			})

			// counterclockwise
			stack = append(stack, &node{
				direction:  n.direction.CounterClockwise(),
				parent:     n,
				moveType:   rotate,
				Coordinate: n.Coordinate,
			})
		}
	}

	// compute the cost of each path!
	minCost := -1
	for _, path := range paths {
		node := path
		cost := 0
		if node.Coordinate != *end {
			panic("how did you end on the wrong place??")
		}
		for node != nil {
			switch node.moveType {
			case step:
				cost++
			case rotate:
				cost += 1000
			}
			node = node.parent
		}
		if minCost == -1 || cost < minCost {
			minCost = cost
		}
	}
	return minCost, nil
}

func isWall(matrix [][]rune, pos shared.Coordinate) bool {
	return matrix[pos.I][pos.J] == '#'
}

func (n *node) previouslySeen(pos shared.Coordinate) bool {
	if n.seen == nil {
		return false
	}
	_, ok := n.seen[pos]
	return ok
}

func (n *node) copySeenAndAdd(pos shared.Coordinate) map[shared.Coordinate]struct{} {
	seen := map[shared.Coordinate]struct{}{}
	for pos := range n.seen {
		seen[pos] = struct{}{}
	}
	seen[pos] = struct{}{}
	return seen
}
