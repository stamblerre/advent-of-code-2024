package main

import (
	"fmt"
	"log"
	"math"
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
	switch part {
	case 1:
		blinks = 25
	case 2:
		blinks = 75
	}
	// stone number -> # times seen
	cache := map[int]int{}
	for _, stone := range stones {
		cache[stone]++
	}
	for range blinks {
		newStones := map[int]int{}
		for stone, count := range cache {
			if stone == 0 {
				newStones[1] += count
			} else if firstHalf, secondHalf, ok := splitDigits(stone); ok {
				newStones[firstHalf] += count
				newStones[secondHalf] += count
			} else {
				newStones[stone*2024] += count

			}
		}
		cache = newStones
	}
	result := 0
	for _, count := range cache {
		result += count
	}
	return result, nil
}

func splitDigits(stone int) (int, int, bool) {
	// apparently you can get the number of digits in a number with floor(1+log(n))
	numDigits := int(math.Floor(1 + math.Log10(float64(stone))))
	if numDigits%2 != 0 {
		return 0, 0, false
	}
	midpoint := numDigits / 2
	firstHalf := stone / int(math.Pow(10, float64(midpoint)))
	secondHalf := stone - firstHalf*int(math.Pow(10, float64(midpoint)))
	return firstHalf, secondHalf, true
}
