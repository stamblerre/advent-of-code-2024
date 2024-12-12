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
	plant     rune
	coords    map[shared.Coordinate]struct{}
	perimeter int // only gets set on calls to mergeRegion?
}

type rowRegion struct {
	plant   rune
	indices []int
}

func implementation(input any, part int) (int, error) {
	matrix, ok := input.([][]rune)
	if !ok {
		return -1, fmt.Errorf("unexpected type for input %T", input)
	}
	var prev []*region
	for i, row := range matrix {
		var curr []*rowRegion
		var currRegion *rowRegion
		for j, plant := range row {
			if currRegion == nil || matrix[i][j] != currRegion.plant {
				currRegion = &rowRegion{
					plant:   plant,
					indices: []int{j},
				}
				curr = append(curr, currRegion)
			} else {
				currRegion.indices = append(currRegion.indices, j)
			}
		}
		merged := mergeRegions(prev, curr, i)
		prev = merged
	}
	result := 0
	for _, r := range prev {
		result += r.perimeter * len(r.coords)
	}
	return result, nil
}

func mergeRegions(original []*region, toAdd []*rowRegion, lastRow int) []*region {
	// you only need to worry about connection to the very last row
	finalizePerimeter := map[*region]struct{}{}
	for _, r := range original {
		foundLastRow := false
		allJs := map[int]struct{}{}
		for coord := range r.coords {
			foundLastRow = foundLastRow || coord.I == lastRow
			if coord.I == lastRow {
				allJs[coord.J] = struct{}{}
			}
		}
		if !foundLastRow {
			// mark this as an umerged region whose perimeter we should finalize
			finalizePerimeter[r] = struct{}{}
			continue
		}
		for index, a := range toAdd {
			if a.plant != r.plant {
				continue
			}
			doMerge := false
			for _, j := range a.indices {
				if _, ok := allJs[j]; ok {
					doMerge = true
					break
				}
			}
			if doMerge {
				for _, j := range a.indices {
					r.coords[shared.Coordinate{
						I: lastRow,
						J: j,
					}] = struct{}{}
				}
				r.perimeter += 2 // you always get left & right for free
				// now count how much overlap there is
				//len(allJs)-len(a.indices)
				// don't try merging with another region...
				toAdd = append(toAdd[0:index], toAdd[index+1:]...)
			}
		}
	}
	// If any row regions remain that we haven't added...convert them to real regions.
	for _, rowReg := range toAdd {
		newReg := &region{
			plant: rowReg.plant,
		}
		for _, j := range rowReg.indices {
			newReg.coords[shared.Coordinate{I: lastRow, J: j}] = struct{}{}
		}
		// perimeter: left + right + top
		newReg.perimeter += 1 + 1 + len(rowReg.indices)
		original = append(original, newReg)
	}
	// for lastRow-1, close up any perimeters that were unmerged
	for r := range finalizePerimeter {
		// get all coordinates for the last row
		count := 0
		for coord := range r.coords {
			if coord.J == lastRow-1 {
				count++
			}
		}
		r.perimeter += count
	}
	return original
}

type node struct {
	*shared.Coordinate
	parent   *node
	children []*node
}

type graphRegion map[shared.Coordinate]struct{}

func graphImplementation(inout any, part int) (int, error) {
	matrix, ok := input.([][]rune)
	if !ok {
		return -1, fmt.Errorf("unexpected type for input %T", input)
	}
	var queue []*node
	for i, row := range matrix {
		for j := range row {
			queue = append(queue, &node{
				Coordinate: &shared.Coordinate{
					I: i,
					J: j,
				},
			})
		}
	}
	seen := map[shared.Coordinate]struct{}{}
	var ends []*node
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		seen[*n.Coordinate] = struct{}{}

		for _, delta := range shared.DirectionDelta() {
			coord := n.Coordinate.Add(&delta)
			if _, ok := seen[*coord]; ok {
				continue
			}
			if !shared.InBounds(matrix, coord) {
				continue
			}
			plant := matrix[coord.I][coord.J]
			if plant == matrix[n.I][n.J] {
				child := &node{
					Coordinate: coord,
					parent:     n,
				}
				n.children = append(n.children, child)
				queue = append(queue, child)
			}
		}
		if len(n.children) == 0 {
			ends = append(ends, n)
		}
	}
	// construct regions
	result := 0
	for _, end := range ends {
		root := end.findRoot()
		r := root.toRegion()
		result += r.area() * r.perimeter(matrix)
	}
	return result, nil
}

func (n *node) findRoot() *node {
	temp := n
	for {
		if temp.parent == nil {
			return temp
		}
		temp = temp.parent
	}
}

func (n *node) toRegion() *graphRegion {
	r := make(graphRegion)

	queue := []*node{n}
	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		r[*next.Coordinate] = struct{}{}

		queue = append(queue, next.children...)
	}
	return &r
}

func (r *graphRegion) area() int {
	return len(*r)
}

func (r *graphRegion) perimeter(matrix [][]rune) int {
	perim := 4*len(*r) - (len(*r) - 1)
	// a side is touching if on the left & right of it there is the same region
	for coord := range *r {
		plant := matrix[coord.I][coord.J]
		for _, delta := range shared.CardinalDirectionDelta() {
			dir := coord.Add(&delta)
			if shared.InBounds(matrix, dir) && matrix[dir.I][dir.J] == plant {
				perim--
			}
		}
	}
	return perim
}
