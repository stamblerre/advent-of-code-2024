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
	intSlice, err := shared.StringSliceToInt(strings.Split(string(text), " "))
	var result []uint64
	for _, val := range intSlice {
		result = append(result, uint64(val))
	}
	return result, nil
}

func (t *today) Part1(input any) (int, error) {
	return implementation(input, 1)
}

func (t *today) Part2(input any) (int, error) {
	return implementation(input, 2)
}

func implementation(input any, part int) (int, error) {
	stones, ok := input.([]uint64)
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
	m := map[int]int{}
	for i := 0; i < blinks; i++ {
		var newStones []uint64
		for _, stone := range stones {
			if stone == 0 {
				newStones = append(newStones, 1)
			} else if firstHalf, secondHalf, ok := splitDigits(stone); ok {
				newStones = append(newStones, firstHalf, secondHalf)
			} else {
				// otherwise...
				newStones = append(newStones, uint64(stone)*2024)
			}
		}
		stones = newStones
		m[i] = len(newStones)
		fmt.Printf("------STONE DELTA %v-------\n", m[i]-m[i-1])
	}
	return len(stones), nil
}

func splitDigits(stone uint64) (uint64, uint64, bool) {
	// apparently you can get the number of digits in a number with floor(1+log(n))
	numDigits := int(math.Floor(1 + math.Log10(float64(stone))))
	if numDigits%2 != 0 {
		return 0, 0, false
	}
	midpoint := numDigits / 2
	firstHalf := stone / uint64(math.Pow(10, float64(midpoint)))
	secondHalf := stone - firstHalf*uint64(math.Pow(10, float64(midpoint)))
	return firstHalf, secondHalf, true
}
