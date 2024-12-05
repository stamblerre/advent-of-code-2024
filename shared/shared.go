package shared

import (
	"os"
	"strconv"
	"strings"
)

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
		i, err := strconv.Atoi(a)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}
