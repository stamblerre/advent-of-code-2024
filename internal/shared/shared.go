package shared

import (
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	I, J int
}

type CoordinateDelta struct {
	DeltaI, DeltaJ int
}

func (c *Coordinate) Add(d *CoordinateDelta) *Coordinate {
	return &Coordinate{
		I: c.I + d.DeltaI,
		J: c.J + d.DeltaJ,
	}
}

func FileToRuneMatrix(filename string) ([][]rune, error) {
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

func RuneSliceString(s []rune) []string {
	var strSlice []string
	for _, s := range s {
		strSlice = append(strSlice, string(s))
	}
	return strSlice
}
