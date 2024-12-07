package main

import (
	"fmt"
	"log"

	"advent-of-code-2024.com/internal/shared"
)

func main() {
	part1, part2, err := run("testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result for part 1: %v\n", part1)
	fmt.Printf("Result for part 2: %v\n", part2)
}

func run(filename string) (int, int, error) {
	input, err := shared.FileToRuneMatrix(filename)
	if err != nil {
		return -1, -1, err
	}
	guard, err := findGuard(input)
	if err != nil {
		return -1, -1, err
	}
	part1, err := mapGuardPath(guard, input)
	if err != nil {
		return -1, -1, err
	}
	part2, err := placeObstacle(guard, input)
	if err != nil {
		return -1, -1, err
	}
	return part1, part2, nil
}

func findGuard(input [][]rune) (*shared.Coordinate, error) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if isGuard(input[i][j]) {
				return &shared.Coordinate{
					I: i,
					J: j,
				}, nil
			}
		}
	}
	return nil, fmt.Errorf("unable to find the guard")

}

func placeObstacle(guard *shared.Coordinate, input [][]rune) (int, error) {
	count := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] != '.' {
				continue
			}
			// test out the obstacle ...
			input[i][j] = '#'
			_, looping, err := patrol(guard, input)
			// and reset...
			input[i][j] = '.'

			if err != nil {
				return -1, err
			}

			if looping {
				count++
			}
		}
	}
	return count, nil
}

func mapGuardPath(guard *shared.Coordinate, input [][]rune) (int, error) {
	visited, _, err := patrol(guard, input)
	return visited, err
}

func patrol(start *shared.Coordinate, input [][]rune) (int, bool, error) {
	loopGuard := map[directionalCoordinate]struct{}{}
	visited := map[shared.Coordinate]struct{}{}
	guard := start
	guardContent := input[guard.I][guard.J]
	for {
		dCoord := directionalCoordinate{
			Coordinate: *guard,
			direction:  guardContent,
		}

		// looping?
		if _, ok := loopGuard[dCoord]; ok {
			return -1, true, nil
		}
		loopGuard[dCoord] = struct{}{}
		visited[*guard] = struct{}{}

		// If there is something directly in front of you, turn right 90 degrees.
		rotations, err := guardRotations(guardContent)
		if err != nil {
			return -1, false, err
		}
		for _, r := range rotations {
			delta := guardToDirection(r)
			next := guard.Add(delta)

			// Trying to go outside the mapped area, so you're done.
			if !inBounds(next, input) {
				return len(visited), false, nil
			}

			if input[next.I][next.J] != '#' {
				guard = next
				guardContent = r
				break
			}
		}
	}
}

func isGuard(r rune) bool {
	switch r {
	case '^', '>', 'v', '<':
		return true
	default:
		return false
	}
}

func guardToDirection(r rune) *shared.CoordinateDelta {
	switch r {
	case '^':
		return &shared.CoordinateDelta{DeltaI: -1, DeltaJ: 0}
	case '>':
		return &shared.CoordinateDelta{DeltaI: 0, DeltaJ: 1}
	case 'v':
		return &shared.CoordinateDelta{DeltaI: 1, DeltaJ: 0}
	case '<':
		return &shared.CoordinateDelta{DeltaI: 0, DeltaJ: -1}
	default:
		return nil
	}
}

func guardRotations(r rune) ([]rune, error) {
	order := []rune{'^', '>', 'v', '<'}
	index := -1
	for i, o := range order {
		if o == r {
			index = i
		}
	}
	if index < 0 {
		return nil, fmt.Errorf("didn't find %s in order", string(r))
	}
	return append(append([]rune{r}, order[index+1:]...), order[:index]...), nil
}

func inBounds(coord *shared.Coordinate, input [][]rune) bool {
	return coord.I >= 0 && coord.J >= 0 && coord.I < len(input) && coord.J < len(input[coord.I])
}

type directionalCoordinate struct {
	shared.Coordinate
	direction rune
}
