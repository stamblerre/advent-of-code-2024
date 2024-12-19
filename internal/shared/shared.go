package shared

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Day interface {
	ReadInput(filename string) (any, error)
	Part1(input any) (int, error)
	Part2(input any) (int, error)
}

func Run(d Day, filename string) (int, int, error) {
	input, err := d.ReadInput(filename)
	if err != nil {
		return -1, -1, err
	}
	part1, err := d.Part1(input)
	if err != nil {
		return -1, -1, err
	}
	part2, err := d.Part2(input)
	if err != nil {
		return -1, -1, err
	}
	fmt.Printf("Result for part 1: %v\n", part1)
	fmt.Printf("Result for part 2: %v\n", part2)
	return part1, part2, nil
}

// random helpful things

type Coordinate struct {
	I, J int
}

func SortedCoordinates(coordinates map[Coordinate]struct{}) []Coordinate {
	var result []Coordinate
	for c := range coordinates {
		result = append(result, c)
	}
	sort.Slice(result, func(a, b int) bool {
		if result[a].I == result[b].I {
			return result[a].J < result[b].J
		}
		return result[a].I < result[b].I
	})
	return result
}

func (c1 *Coordinate) Neighbors(c2 *Coordinate) bool {
	for _, d := range CardinalDirectionDelta() {
		neighbor := c1.Add(&d)
		if neighbor.equals(c2) {
			return true
		}
	}
	return false
}

func (c1 *Coordinate) equals(c2 *Coordinate) bool {
	return c1.I == c2.I && c1.J == c2.J
}

type CoordinateDelta struct {
	DeltaI, DeltaJ int
}

func (d *CoordinateDelta) Multiply(i int) *CoordinateDelta {
	return &CoordinateDelta{
		DeltaI: d.DeltaI * i,
		DeltaJ: d.DeltaJ * i,
	}
}

func (c *Coordinate) Add(d *CoordinateDelta) *Coordinate {
	return &Coordinate{
		I: c.I + d.DeltaI,
		J: c.J + d.DeltaJ,
	}
}

func (c *Coordinate) Sub(d *CoordinateDelta) *Coordinate {
	return &Coordinate{
		I: c.I - d.DeltaI,
		J: c.J - d.DeltaJ,
	}
}

func (c1 *Coordinate) Delta(c2 *Coordinate) *CoordinateDelta {
	return &CoordinateDelta{
		DeltaI: c1.I - c2.I,
		DeltaJ: c1.J - c2.J,
	}
}

func InBounds(matrix any, coord *Coordinate) bool {
	switch matrix := matrix.(type) {
	case [][]rune:
		return coord.I >= 0 && coord.J >= 0 && coord.I < len(matrix) && coord.J < len(matrix[coord.I])
	case [][]int:
		return coord.I >= 0 && coord.J >= 0 && coord.I < len(matrix) && coord.J < len(matrix[coord.I])
	default:
		panic(fmt.Sprintf("unexpected type %T for matrix", matrix))
	}
}

type Direction int

const (
	Right Direction = iota
	Down
	Left
	Up
	DiagonalUpRight
	DiagonalUpLeft
	DiagonalDownRight
	DiagonalDownLeft
)

// TODO(stamblerre): redo this

func CardinalDirectionDelta() map[Direction]CoordinateDelta {
	return map[Direction]CoordinateDelta{
		Right: {DeltaI: 0, DeltaJ: 1},
		Down:  {DeltaI: 1, DeltaJ: 0},
		Left:  {DeltaI: 0, DeltaJ: -1},
		Up:    {DeltaI: -1, DeltaJ: 0},
	}
}

func DirectionDelta() map[Direction]CoordinateDelta {
	m := CardinalDirectionDelta()

	m[DiagonalUpRight] = CoordinateDelta{DeltaI: 1, DeltaJ: 1}
	m[DiagonalDownRight] = CoordinateDelta{DeltaI: 1, DeltaJ: -1}
	m[DiagonalDownLeft] = CoordinateDelta{DeltaI: -1, DeltaJ: 1}
	m[DiagonalUpLeft] = CoordinateDelta{DeltaI: -1, DeltaJ: -1}

	return m
}

// reading files

func ReadRuneMatrix(filename string) ([][]rune, error) {
	text, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	// convert to 2D array
	var input [][]rune
	for i, line := range strings.Split(string(text), "\n") {
		input = append(input, []rune{})
		for _, c := range line {
			input[i] = append(input[i], c)
		}
	}
	return input, nil
}

func ReadIntMatrix(filename string) ([][]int, error) {
	runes, err := ReadRuneMatrix(filename)
	if err != nil {
		return nil, err
	}
	var result [][]int
	for _, line := range runes {
		asInts, err := RuneSliceToInt(line)
		if err != nil {
			return nil, err
		}
		result = append(result, asInts)
	}
	return result, nil
}

func StringSliceToInt(strOfInts []string) ([]int, error) {
	var slice []any
	for _, s := range strOfInts {
		slice = append(slice, s)
	}
	return sliceToInt(slice)
}

func RuneSliceToInt(strOfInts []rune) ([]int, error) {
	var slice []any
	for _, s := range strOfInts {
		slice = append(slice, s)
	}
	return sliceToInt(slice)
}

func sliceToInt(strOfInts []any) ([]int, error) {
	var result []int
	for _, a := range strOfInts {
		if a == "" {
			continue
		}
		var s string
		switch a := a.(type) {
		case rune:
			s = string(a)
		case string:
			s = a
		default:
			return nil, fmt.Errorf("unexpected type %T for int conversion", a)
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

func PrintRuneMatrix(matrix [][]rune) {
	for _, line := range matrix {
		fmt.Println(RuneSliceString(line))
	}
}

func RuneSliceString(s []rune) []string {
	var strSlice []string
	for _, s := range s {
		strSlice = append(strSlice, string(s))
	}
	return strSlice
}

type Point struct {
	X, Y int
}
