package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	text, err := os.ReadFile("testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	list1, list2, err := getSortedSlices(string(text))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Min distance is: %v\n", getMinDistance(list1, list2))
	fmt.Printf("Similarity score: %v\n", getSimilarityScore(list1, list2))
}

func getSimilarityScore(left, right []int) int {
	rightCount := map[int]int{}
	for _, val := range right {
		rightCount[val]++
	}
	similarityScore := 0
	for _, val := range left {
		similarityScore += val * rightCount[val]
	}
	return similarityScore
}

func getMinDistance(list1, list2 []int) int {
	sort.Ints(list1)
	sort.Ints(list2)

	var difference float64
	for i, val1 := range list1 {
		val2 := list2[i]
		difference += math.Abs(float64(val2 - val1))
	}
	return int(difference)
}

func getSortedSlices(input string) ([]int, []int, error) {
	var list1, list2 []int
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, "   ")
		if len(split) != 2 {
			return nil, nil, fmt.Errorf("expected 2 items, got %s", line)
		}
		val1, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert %s", split[0])
		}
		val2, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert %s", split[1])
		}
		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	return list1, list2, nil
}
