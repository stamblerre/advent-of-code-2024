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

type region struct {
	coordinates map[shared.Coordinate]struct{}
	matrix      *[][]rune
}

func implementation(input any, part int) (int, error) {
	matrix, ok := input.([][]rune)
	if !ok {
		return -1, fmt.Errorf("unexpected type for input %T", input)
	}
	// grab all regions by letter
	letterRegions := map[rune]*region{}
	for i, row := range matrix {
		for j := range row {
			plant := matrix[i][j]
			if _, ok := letterRegions[plant]; !ok {
				letterRegions[plant] = &region{
					coordinates: make(map[shared.Coordinate]struct{}),
					matrix:      &matrix,
				}
			}
			letterRegions[plant].coordinates[shared.Coordinate{
				I: i,
				J: j,
			}] = struct{}{}
		}
	}
	// check if we need to split any regions up
	var knownRegions []*region
	for _, region := range letterRegions {
		split := splitRegions(region)
		knownRegions = append(knownRegions, split...)
	}
	result := 0
	for _, r := range knownRegions {
		switch part {
		case 1:
			result += r.perimeter() * r.area()
		case 2:
			result += r.sides() * r.area()
		default:
			panic(fmt.Sprintf("unexpected part %v", part))
		}

	}
	return result, nil
}

func splitRegions(originalRegion *region) []*region {
	sorted := shared.SortedCoordinates(originalRegion.coordinates)
	currentRegion := &region{
		coordinates: map[shared.Coordinate]struct{}{},
		matrix:      originalRegion.matrix,
	}
	var regions []*region
	for len(sorted) > 0 {
		index := currentRegion.getNextMergeable(sorted)
		if index == -1 {
			regions = append(regions, currentRegion)
			currentRegion = &region{
				coordinates: map[shared.Coordinate]struct{}{},
				matrix:      originalRegion.matrix,
			}
			continue
		}
		coord := sorted[index]
		sorted = append(sorted[0:index], sorted[index+1:]...)
		currentRegion.coordinates[coord] = struct{}{}
	}
	if len(currentRegion.coordinates) > 0 {
		regions = append(regions, currentRegion)
	}
	return regions
}

func (r *region) getNextMergeable(coords []shared.Coordinate) int {
	if len(r.coordinates) == 0 {
		return 0
	}
	for i, coord := range coords {
		if r.mergeable(coord) {
			return i
		}
	}
	return -1
}

func (r *region) mergeable(coord shared.Coordinate) bool {
	if len(r.coordinates) == 0 {
		return true
	}
	for c := range r.coordinates {
		if c.Neighbors(&coord) {
			return true
		}
	}
	return false
}

func (r *region) area() int {
	return len(r.coordinates)
}

func (r *region) plant() string {
	for coord := range r.coordinates {
		plant := (*r.matrix)[coord.I][coord.J]
		return string(plant)
	}
	return "NO PLANT"
}

func (r *region) perimeter() int {
	perim := 4 * len(r.coordinates)
	// a side is touching if on the left & right of it there is the same region
	for coord := range r.coordinates {
		plant := (*r.matrix)[coord.I][coord.J]
		for _, delta := range shared.CardinalDirectionDelta() {
			dir := coord.Add(delta)
			if !shared.InBounds(*r.matrix, *dir) {
				continue
			}
			dirPlant := (*r.matrix)[dir.I][dir.J]
			if dirPlant == plant {
				perim -= 1
			}
		}
	}
	return perim
}

func (r *region) sides() int {
	perim := 4 * len(r.coordinates)
	// a side is touching if on the left & right of it there is the same region
	for coord := range r.coordinates {
		plant := (*r.matrix)[coord.I][coord.J]
		for _, delta := range shared.CardinalDirectionDelta() {
			dir := coord.Add(delta)
			if !shared.InBounds(*r.matrix, *dir) {
				continue
			}
			dirPlant := (*r.matrix)[dir.I][dir.J]
			if dirPlant == plant {
				perim -= 1
			}
		}
	}
	return perim
}
