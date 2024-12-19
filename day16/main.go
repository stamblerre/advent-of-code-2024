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

	// cost of the move...
	moveCost int
}

func (n *node) String() string {
	return fmt.Sprintf("%v - %s", n.Coordinate, n.direction)
}

const (
	stepCost     = 1
	rotationCost = 1000
)

func implementation(input any, part int) (int, error) {
	matrix, ok := input.([][]rune)
	if !ok {
		return -1, fmt.Errorf("unexpected type %T for input", input)
	}
	start := shared.FindRune(matrix, 'S')
	end := shared.FindRune(matrix, 'E')

	queue := []*node{{
		Coordinate: *start,
		direction:  shared.Right,
	}}
	// i -> j -> min cost
	cost := map[int]map[int]int{
		start.I: {start.J: 0},
	}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]

		// check that you can walk here
		value := matrix[n.I][n.J]
		if value == '#' {
			continue
		}

		// update the cost
		if _, ok := cost[n.I]; !ok {
			cost[n.I] = map[int]int{}
		}
		var newCost int
		if n.parent != nil {
			newCost += cost[n.parent.I][n.parent.J]
		}
		newCost += n.moveCost

		if minCost, ok := cost[n.I][n.J]; !ok || newCost < minCost {
			cost[n.I][n.J] = newCost
		}

		// reached the end!
		if n.Coordinate.Equals(end) {
			continue
		}

		// options:
		//
		// (1) go forward, if possible
		if next := n.AddDirection(n.direction); shared.InBounds(matrix, next) && !isWall(matrix, next) && !n.previouslySeen(next) {
			queue = append(queue, &node{
				direction:  n.direction,
				parent:     n,
				moveCost:   stepCost,
				Coordinate: n.AddDirection(n.direction),
				seen:       n.copySeenAndAdd(next),
			})
		}

		// (2) turn clockwise or counterclockwise
		// but don't turn if you've turned before...don't want to spin in circles
		if n.moveCost <= stepCost {
			// clockwise
			for _, newDir := range []shared.Direction{
				n.direction.Clockwise(),
				n.direction.CounterClockwise(),
			} {
				newCoord := n.Coordinate.AddDirection(newDir)
				if shared.InBounds(matrix, newCoord) {
					queue = append(queue, &node{
						direction:  newDir,
						parent:     n,
						moveCost:   rotationCost + stepCost,
						Coordinate: newCoord,
						seen:       n.copySeenAndAdd(newCoord),
					})
				}
			}
		}
	}
	return cost[end.I][end.J], nil
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
