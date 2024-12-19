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
		fmt.Printf("---------------MOVE: %v------------------\n", string(move))
		dir := shared.CaratToDirection(move)
		if dir == shared.Unknown {
			panic(fmt.Sprintf("unknown direction %v", move))
		}
		delta := shared.DirectionToDelta(dir)
		nextPos := pos.Add(delta)
		next := in.grid[nextPos.I][nextPos.J]
		switch next {
		case 'O':
			// push boxes
			if pushed := pushBoxes(&in.grid, nextPos, delta); pushed {
				fmt.Println("PUSHED!")
				pos = nextPos
			} else {
				fmt.Println("I HATE TO PUSH")
			}
		case '#':
			// do nothing
			fmt.Println("nothing")
		case '.':
			// move the robot
			in.grid[pos.I][pos.J] = '.'
			in.grid[nextPos.I][nextPos.J] = '@'
			pos = nextPos
			fmt.Println("move da robot")
		}
		shared.PrintRuneMatrix(in.grid)
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
	for i, line := range grid {
		for j, r := range line {
			if r == '@' {
				return &shared.Coordinate{
					I: i,
					J: j,
				}
			}
		}
	}
	return nil
}

func pushBoxes(grid *[][]rune, boxPos *shared.Coordinate, delta shared.CoordinateDelta) bool {
	fmt.Println("Pushing boxes...")
	var empty *shared.Coordinate
	pos := boxPos
	for {
		next := pos.Add(delta)
		if !shared.InBounds(*grid, next) {
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
		} else {
			panic(fmt.Sprintf("unexpected value %v", string(value)))
		}
	}
	// no space to push, just return
	if empty == nil {
		return false
	}
	fmt.Printf("I FOUND AN EMPTY PLACE %v\n", empty)
	pos = empty
	for pos != boxPos {
		prevPos := pos.Sub(delta)
		if !shared.InBounds(*grid, prevPos) {
			break
		}
		prevPosValue := (*grid)[prevPos.I][prevPos.J]
		if prevPosValue == '#' {
			break // don't move the walls
		}
		fmt.Printf("PREV POS VALUE %v\n", string(prevPosValue))
		posValue := (*grid)[pos.I][pos.J]
		(*grid)[pos.I][pos.J] = prevPosValue
		(*grid)[prevPos.I][prevPos.J] = posValue
		pos = prevPos
	}
	return true
}
