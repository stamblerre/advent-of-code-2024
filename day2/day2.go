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
		increasing := false
		isFirstViolation := true
		for i, val := range report {
			if i == 0 {
				continue
			}

			prev := report[i-1]
			// check if it's increasing or decreasing at the first value
			if i == 1 {
				if prev < val {
					increasing = true
				}
			}

			hasViolation := prev == val || (increasing && (val > prev+3 || val < prev)) || (!increasing && (val < prev-3 || val > prev))
			if hasViolation {
				// this doesnt work
				if useDampener && isFirstViolation {
					isFirstViolation = false
					continue
				}
				// in all other cases, record the violation
				break
			}

			// we got to the end without breaking out of the loop
			if i == len(report)-1 {
				countSafe++
			}
		}
	}
	return countSafe, nil
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
