package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strings"

	"advent-of-code-2024.com/shared"
)

func main() {
	part1, part2, err := validateUpdates("testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result for part 1: %v\n", part1)
	fmt.Printf("Result for part 2: %v\n", part2)
}

func validateUpdates(filename string) (int, int, error) {
	orders, updates, err := readInput(filename)
	if err != nil {
		return -1, -1, err
	}
	validResult := 0
	fixedResult := 0
	for _, after := range updates {
		before := make([]int, len(after))
		copy(before, after)

		sort.SliceStable(after, func(i, j int) bool {
			return evaluateOrdering(orders, after[i], after[j])
		})
		midpoint := after[len(after)/2]
		if slices.Equal(before, after) {
			validResult += midpoint
		} else {
			fixedResult += midpoint
		}
	}
	return validResult, fixedResult, nil
}

func evaluateOrdering(rules []pageOrder, page1, page2 int) bool {
	for _, rule := range rules {
		if rule.x == page1 && rule.y == page2 {
			return true
		}
	}
	return false
}

type pageOrder struct {
	x, y int
}

func readInput(filename string) ([]pageOrder, [][]int, error) {
	text, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}
	var orders []pageOrder
	var updates [][]int
	for _, line := range strings.Split(string(text), "\n") {
		if order := strings.Split(line, "|"); len(order) == 2 {
			orderAsInts, err := shared.StrSliceToInt(order)
			if err != nil {
				return nil, nil, err
			}
			orders = append(orders, pageOrder{x: orderAsInts[0], y: orderAsInts[1]})
		} else if update := strings.Split(line, ","); len(update) > 1 {
			updateAsInt, err := shared.StrSliceToInt(update)
			if err != nil {
				return nil, nil, err
			}
			updates = append(updates, updateAsInt)
		}
	}
	return orders, updates, nil
}
