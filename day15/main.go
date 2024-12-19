package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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

type gridAndMoves struct {
	grid  [][]rune
	moves []rune
}

func (t *today) ReadInput(filename string) (any, error) {
	text, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(text), "\n\n")
	grid, err := shared.TextToRuneMatrix(split[0])
	if err != nil {
		return nil, err
	}
	return &gridAndMoves{
		grid:  grid,
		moves: []rune(strings.Replace(split[1], "\n", "", -1)),
	}, nil
}

func (t *today) Part1(input any) (int, error) {
	return implementation(input, 1)
}

func (t *today) Part2(input any) (int, error) {
	return implementation(input, 2)
}

func implementation(input any, part int) (int, error) {
	in, ok := input.(*gridAndMoves)
	if !ok {
		return -1, fmt.Errorf("unexpected type %T for input", input)
	}
	// find the robot
	pos := findRobot(in.grid)
	if pos == nil {
		return -1, fmt.Errorf("unable to find start of robot")
	}
	for _, move := range in.moves {
		dir := shared.CaratToDirection(move)
		if dir == shared.Unknown {
			panic(fmt.Sprintf("unknown direction %v", move))
		}
		delta := shared.DirectionToDelta(dir)
		nextPos := pos.Add(delta)
		next := in.grid[nextPos.I][nextPos.J]
		moveRobot := func() {
			in.grid[pos.I][pos.J] = '.'
			in.grid[nextPos.I][nextPos.J] = '@'
			pos = nextPos
		}
		switch next {
		case 'O':
			// push boxes
			if pushed := pushBoxes(&in.grid, nextPos, delta); pushed {
				moveRobot()
			}
		case '#':
			// do nothing
		case '.':
			// move the robot
			moveRobot()
		}
		if countRobots(in.grid) > 1 {
			os.Exit(1)
		}
	}
	gps := 0
	for i, line := range in.grid {
		for j, r := range line {
			// not a box
			if r != 'O' {
				continue
			}
			gps += 100*i + j
		}
	}
	return gps, nil
}

func findRobot(grid [][]rune) *shared.Coordinate {
	return shared.FindRune(grid, '@')
}

func pushBoxes(grid *[][]rune, newRobotPos *shared.Coordinate, delta shared.CoordinateDelta) bool {
	var empty *shared.Coordinate
	pos := newRobotPos
	for {
		next := pos.Add(delta)
		if !shared.InBounds(*grid, *next) {
			break
		}
		value := (*grid)[next.I][next.J]
		if value == '#' {
			// reached the end...
			break
		} else if value == '.' {
			//  found an empty spot
			empty = next
			break
		} else if value == 'O' {
			pos = next
		} else if value == '@' {
			panic(fmt.Sprintf("found robot at %v, newRobotPos is %v", next, newRobotPos))
		} else {
			panic(fmt.Sprintf("unexpected value %v", string(value)))
		}
	}
	// no space to push, just return
	if empty == nil {
		return false
	}
	pos = empty
	for !pos.Equals(newRobotPos) {
		prevPos := pos.Sub(delta)
		if !shared.InBounds(*grid, *prevPos) {
			break
		}
		prevPosValue := (*grid)[prevPos.I][prevPos.J]
		if prevPosValue == '#' {
			break // don't move the walls
		}
		if prevPosValue == '@' {
			// i saw the robot...?
			panic(fmt.Sprintf("found the robot at %v", prevPosValue))
		}
		posValue := (*grid)[pos.I][pos.J]
		if posValue == '@' {
			// i saw the robot...?
			panic(fmt.Sprintf("found the robot at %v", posValue))
		}
		(*grid)[pos.I][pos.J] = prevPosValue
		(*grid)[prevPos.I][prevPos.J] = posValue
		pos = prevPos
	}

	return true
}

func countRobots(grid [][]rune) int {
	seen := 0
	for _, line := range grid {
		for _, r := range line {
			if r == '@' {
				seen++
			}
		}
	}
	return seen
}
