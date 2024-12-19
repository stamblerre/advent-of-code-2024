package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

type today struct{}

var (
	inputRe = regexp.MustCompile(`p=(\d+),(\d+) v=(-*\d+),(-*\d+)`)
)

type robot struct {
	position *shared.Coordinate
	velocity shared.CoordinateDelta
}

func (t *today) ReadInput(filename string) (any, error) {
	text, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var robots []*robot
	for _, line := range strings.Split(string(text), "\n") {
		matches := inputRe.FindStringSubmatch(line)
		if len(matches) != 5 {
			panic(fmt.Sprintf("not sure what to do with line: %q", line))
		}
		intSlice, err := shared.StringSliceToInt(matches[1:])
		if err != nil {
			return nil, err
		}
		robots = append(robots, &robot{
			position: &shared.Coordinate{
				J: intSlice[0], // J = X
				I: intSlice[1], // I = Y
			},
			velocity: shared.CoordinateDelta{
				DeltaJ: intSlice[2], // J = X
				DeltaI: intSlice[3], // I = Y
			},
		})
	}
	return robots, nil
}

func (t *today) Part1(input any) (int, error) {
	return implementation(input, 1)
}

func (t *today) Part2(input any) (int, error) {
	return implementation(input, 2)
}

func implementation(input any, part int) (int, error) {
	if part == 2 {
		return -1, nil
	}
	robots, ok := input.([]*robot)
	if !ok {
		return -1, fmt.Errorf("unexpected type %T for input", input)
	}
	maxX, maxY := robotGridSize(robots)
	seconds := 100
	for _, r := range robots {
		r.move(seconds, maxX, maxY)
	}
	quadrantMap := map[quadrant]int{}
	for _, r := range robots {
		quadrantMap[r.quadrant(maxX, maxY)]++
	}
	result := 1
	for quadrant, factor := range quadrantMap {
		if quadrant == none {
			continue
		}
		result *= factor
	}
	return result, nil
}

func (r *robot) move(seconds, maxX, maxY int) {
	delta := r.velocity.Multiply(seconds)
	newPosition := r.position.Add(*delta)
	for newPosition.I >= maxY {
		newPosition = newPosition.Add(shared.CoordinateDelta{
			DeltaJ: 0,
			DeltaI: -maxY,
		})
	}
	for newPosition.I < 0 {
		newPosition = newPosition.Add(shared.CoordinateDelta{
			DeltaJ: 0,
			DeltaI: maxY,
		})
	}
	for newPosition.J >= maxX {
		newPosition = newPosition.Add(shared.CoordinateDelta{
			DeltaJ: -maxX,
			DeltaI: 0,
		})
	}
	for newPosition.J < 0 {
		newPosition = newPosition.Add(shared.CoordinateDelta{
			DeltaJ: maxX,
			DeltaI: 0,
		})
	}
	r.position = newPosition
}

type quadrant int

const (
	none quadrant = iota
	topLeft
	topRight
	bottomRight
	bottomLeft
)

func (r *robot) quadrant(maxX, maxY int) quadrant {
	if r.position.I > maxY/2 && r.position.J > maxX/2 {
		return topRight
	} else if r.position.I > maxY/2 && r.position.J < maxX/2 {
		return topLeft
	} else if r.position.I < maxY/2 && r.position.J > maxX/2 {
		return bottomRight
	} else if r.position.I < maxY/2 && r.position.J < maxX/2 {
		return bottomLeft
	} else {
		return none
	}
}

func robotGridSize(robots []*robot) (int, int) {
	maxX := -1
	maxY := -1
	for _, r := range robots {
		if maxX == -1 || r.position.J > maxX {
			maxX = r.position.J
		}
		if maxY == -1 || r.position.I > maxY {
			maxY = r.position.I
		}
	}
	return maxX + 1, maxY + 1
}

func robotsToGrid(robots []*robot) string {
	grid := robotsToGridHelper(robots)
	return shared.StringOfStringMatrix(grid)
}

func robotsToGridQuadrant(robots []*robot) string {
	grid := robotsToGridHelper(robots)
	midY := len(grid) / 2
	midX := len(grid[0]) / 2
	for i := range len(grid) {
		grid[i][midX] = ""
	}
	for i := range len(grid[0]) {
		grid[midY][i] = ""
	}
	return shared.StringOfStringMatrix(grid)
}

func robotsToGridHelper(robots []*robot) [][]string {
	seen := map[shared.Coordinate]int{}
	for _, r := range robots {
		seen[*r.position]++
	}
	maxX, maxY := robotGridSize(robots)
	result := make([][]string, maxY)
	for i := range maxY {
		result[i] = make([]string, maxX)
		for j := range maxX {
			if value, ok := seen[shared.Coordinate{
				I: i,
				J: j,
			}]; ok {
				result[i][j] = fmt.Sprintf("%d", value)
			} else {
				result[i][j] = "."
			}
		}
	}
	return result
}
