package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	result, err := wordSearch("testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Count: %v\n", result)
}

func wordSearch(filename string) (int, error) {
	text, err := os.ReadFile(filename)
	if err != nil {
		return -1, err
	}
	// convert to 2D array
	var input [][]rune
	for i, line := range strings.Split(string(text), "\n") {
		input = append(input, []rune{})
		for _, c := range line {
			input[i] = append(input[i], c)
		}
	}
	transposedInput := transposeVertical(input)
	linearCount := countLinear(input) + countLinear(transposedInput)
	fmt.Printf("Linear count: %v\n", linearCount)
	horizontalFlip := flipHorizontally(input)
	fmt.Println(toString(horizontalFlip))
	verticalFlip := flipVertically(input)
	fmt.Println(toString(verticalFlip))
	diagonalCount := countDiagonal(input) + countDiagonal(transposedInput) + countDiagonal(horizontalFlip) + countDiagonal(verticalFlip)
	return linearCount + diagonalCount, nil
}

func toString(input [][]rune) string {
	var result string
	for _, line := range input {
		result += string(line) + "\n"
	}
	return result
}

func countLinear(input [][]rune) int {
	count := 0
	for _, line := range input {
		// first forwards
		count += containsXmasLinear(line)

		// then reverse & try backwards
		reversed := []rune{}
		reversed = append(reversed, line...)
		slices.Reverse(reversed)
		count += containsXmasLinear(reversed)
	}
	return count
}

func transposeVertical(input [][]rune) [][]rune {
	// assumes that every line has the same length...
	verticalInput := make([][]rune, len(input[0]))
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if verticalInput[j] == nil {
				verticalInput[j] = make([]rune, len(input))
			}
			verticalInput[j][i] = input[i][j]
		}
	}
	return verticalInput
}

func flipHorizontally(input [][]rune) [][]rune {
	flipped := make([][]rune, len(input))
	for i, line := range input {
		flipped[i] = make([]rune, len(line))
		for j := 0; j < len(line); j++ {
			flipped[i][len(line)-j-1] = line[j]
		}
	}
	return flipped
}

func flipVertically(input [][]rune) [][]rune {
	flipped := make([][]rune, len(input))
	for i, line := range input {
		row := len(input) - 1 - i
		flipped[row] = make([]rune, len(line))
		copy(flipped[row], line)
	}
	return flipped
}

func containsXmasLinear(line []rune) int {
	count := 0
	for i := 0; i < len(line); i++ {
		indexMap := map[rune]int{
			'X': i,
			'M': i + 1,
			'A': i + 2,
			'S': i + 3,
		}
		if indexMap['S'] >= len(line) {
			continue
		}
		valid := true
		for r, index := range indexMap {
			if line[index] != r {
				valid = false
				break
			}
		}
		if valid {
			count++
		}
	}
	return count
}

func countDiagonal(input [][]rune) int {
	count := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			rowIndexMap := map[rune]int{
				'X': i,
				'M': i + 1,
				'A': i + 2,
				'S': i + 3,
			}
			columnIndexMap := map[rune]int{
				'X': j,
				'M': j + 1,
				'A': j + 2,
				'S': j + 3,
			}
			if rowIndexMap['S'] >= len(input) {
				continue
			}
			if columnIndexMap['S'] >= len(input[i]) {
				continue
			}
			valid := true
			for r, rowIndex := range rowIndexMap {
				colIndex := columnIndexMap[r]
				if input[rowIndex][colIndex] != r {
					valid = false
					break
				}
			}
			if valid {
				count++
			}
		}
	}
	return count
}
