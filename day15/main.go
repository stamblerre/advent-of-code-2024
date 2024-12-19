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
		next := in.grid[pos.I][pos.J]
		switch next {
		case 'O':
			// push boxes
			if pushed := pushBoxes(in.grid, nextPos, dir); pushed {
				pos = nextPos
			}
		case '#':
			// do nothing
		case '.':
			pos = nextPos
		}
		shared.PrintRuneMatrix(in.grid)
		fmt.Printf("---------------------------------\n")
	}
	return -1, nil
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

func pushBoxes(grid [][]rune, boxPos *shared.Coordinate, dir shared.Direction) bool {
	return false
}
