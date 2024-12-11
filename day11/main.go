package main

import (
	"fmt"
	"log"
	"os"
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

type today struct {
}

func (t *today) ReadInput(filename string) (any, error) {
	text, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return shared.StringSliceToInt(strings.Split(string(text), " "))
}

func (t *today) Part1(input any) (int, error) {
	return implementation(input, 1)
}

func (t *today) Part2(input any) (int, error) {
	return implementation(input, 2)
}

func implementation(input any, part int) (int, error) {
	stones, ok := input.([]int)
	if !ok {
		return -1, fmt.Errorf("unexpected type for input %T", input)
	}
	var blinks int
	if part == 1 {
		blinks = 2
	}
	for range blinks {
		var newStones []int
		for _, stone := range stones {
			// apparently you can get the number of digits in a number with floor(1+log(n))
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if firstHalf, secondHalf, ok := splitDigits(stone); ok {
				newStones = append(newStones, firstHalf, secondHalf)
			} else {
				// otherwise...
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
	}
	return len(stones), nil
}

func splitDigits(value int) (int, int, bool) {
	fmt.Printf("NUM DIGITS %v STONE %v\n", numDigits, stone)
}
