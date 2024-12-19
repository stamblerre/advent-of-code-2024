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

type antenna struct {
	shared.Coordinate

	frequency rune
}

func (t *today) Part1(input any) (int, error) {
	return helper(input /*includeHarmonics=*/, false)
}

func (t *today) Part2(input any) (int, error) {
	return helper(input /*includeHarmonics=*/, true)
}

func helper(input any, includeHarmonics bool) (int, error) {
	matrix, ok := input.([][]rune)
	if !ok {
		return -1, fmt.Errorf("unexpected type %T of input", input)
	}
	var antennas []*antenna
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '.' {
				continue
			}
			antennas = append(antennas, &antenna{
				Coordinate: shared.Coordinate{
					I: i,
					J: j,
				},
				frequency: matrix[i][j],
			})
		}
	}
	// create the lines of all the antennas
	antinodes := map[shared.Coordinate]struct{}{}
	for i, antenna1 := range antennas {
		for j, antenna2 := range antennas {
			// don't compare the same antenna
			if i == j {
				continue
			}
			// don't compare antennas of different frequency
			if antenna1.frequency != antenna2.frequency {
				continue
			}

			if !includeHarmonics {
				// should be 2 possible antinodes, 1 where antenna1
				// is 2X away, another where antenna2 is 2X away
				delta := antenna1.Delta(&antenna2.Coordinate).Multiply(2)

				for i, antenna := range []*antenna{antenna1, antenna2} {
					var antinode *shared.Coordinate
					if i == 0 {
						antinode = antenna.Sub(*delta)
					} else {
						antinode = antenna.Add(*delta)
					}
					if !shared.InBounds(matrix, antinode) {
						continue
					}
					antinodes[*antinode] = struct{}{}
				}
			} else {
				// can be way more antinodes now
				// start by subtracting from antenna1
				delta := antenna1.Delta(&antenna2.Coordinate)
				antinode := &antenna1.Coordinate
				for {
					antinode = antinode.Sub(*delta)
					if !shared.InBounds(matrix, antinode) {
						break
					}
					antinodes[*antinode] = struct{}{}
				}
				// now add to antenna2
				antinode = &antenna2.Coordinate
				for {
					antinode = antinode.Add(*delta)
					if !shared.InBounds(matrix, antinode) {
						break
					}
					antinodes[*antinode] = struct{}{}
				}
			}
		}
	}
	return len(antinodes), nil
}
