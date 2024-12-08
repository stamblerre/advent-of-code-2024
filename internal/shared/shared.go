package shared

import (
	"fmt"
	"os"
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

func InBounds(matrix [][]rune, coord *Coordinate) bool {
	return coord.I >= 0 && coord.J >= 0 && coord.I < len(matrix) && coord.J < len(matrix[coord.I])
}

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

func StrSliceToInt(strOfInts []string) ([]int, error) {
	var result []int
	for _, a := range strOfInts {
		if a == "" {
			continue
		}
		i, err := strconv.Atoi(a)
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
