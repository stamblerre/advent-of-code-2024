package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	numSafeReports, err := countSafeReports("testdata/input.txt", true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Num safe levels: %v\n", numSafeReports)
}

func countSafeReports(filename string, useDampener bool) (int, error) {
	reports, err := readReports(filename)
	if err != nil {
		return -1, err
	}
	countSafe := 0
	for _, report := range reports {
		if isSafe(report) {
			countSafe++
			continue
		}
		if useDampener && isSafeWithDampener(report) {
			countSafe++
		}
	}
	return countSafe, nil
}

func isSafeWithDampener(report []int) bool {
	variants := [][]int{}
	for i := 0; i < len(report); i++ {
		variants = append(variants, append(append([]int{}, report[0:i]...), report[i+1:]...))
	}
	for _, variant := range variants {
		if isSafe(variant) {
			return true
		}
	}
	return false
}

func isSafe(report []int) bool {
	increasing := false
	for i, val := range report {
		if i == 0 {
			continue
		}

		prev := report[i-1]
		if i == 1 && prev < val {
			increasing = true
		}
		if prev == val || (increasing && (val > prev+3 || val < prev)) || (!increasing && (val < prev-3 || val > prev)) {
			return false
		}
	}
	return true
}

func readReports(filename string) ([][]int, error) {
	text, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var result [][]int
	for i, line := range strings.Split(string(text), "\n") {
		result = append(result, []int{})
		for _, val := range strings.Split(line, " ") {
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			result[i] = append(result[i], intVal)
		}
	}
	return result, nil
}
