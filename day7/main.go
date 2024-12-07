package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"advent-of-code-2024.com/internal/shared"
)

func main() {
	part1, part2, err := run("testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result for part 1: %v\n", part1)
	fmt.Printf("Result for part 2: %v\n", part2)
}

func run(filename string) (int, int, error) {
	input, err := readInput(filename)
	if err != nil {
		return -1, -1, err
	}
	part1, err := part1(input)
	if err != nil {
		return -1, -1, err
	}
	part2, err := part2(input)
	if err != nil {
		return -1, -1, err
	}
	return part1, part2, nil
}

type equation struct {
	testValue int
	inputs    []int
}

func part1(equations []*equation) (int, error) {
	result := 0
	for _, e := range equations {
		solveable, err := solveable(e /*useConcatenate*/, false)
		if err != nil {
			return -1, err
		}
		if solveable {
			result += e.testValue
		}
	}
	return result, nil
}

func part2(equations []*equation) (int, error) {
	result := 0
	for _, e := range equations {
		solveable, err := solveable(e /*useConcatenate*/, true)
		if err != nil {
			return -1, err
		}
		if solveable {
			result += e.testValue
		}
	}
	return result, nil
}

func solveable(e *equation, useConcatenate bool) (bool, error) {
	acc, err := findAllCombinations(e.inputs, nil, useConcatenate)
	if err != nil {
		return false, err
	}
	for _, val := range acc {
		if val == e.testValue {
			return true, nil
		}
	}
	return false, nil
}

func findAllCombinations(inputs []int, acc []int, useConcatenate bool) ([]int, error) {
	if len(inputs) == 0 {
		return acc, nil
	}
	if acc == nil {
		acc = []int{inputs[0]}
		inputs = inputs[1:]
	}
	var newAcc []int
	for _, a := range acc {
		newAcc = append(newAcc, a*inputs[0])
		newAcc = append(newAcc, a+inputs[0])
		if useConcatenate {
			concatenated, err := concatenate(a, inputs[0])
			if err != nil {
				return nil, err
			}
			newAcc = append(newAcc, concatenated)
		}
	}
	return findAllCombinations(inputs[1:], newAcc, useConcatenate)
}

func concatenate(a, b int) (int, error) {
	combined := fmt.Sprintf("%d%d", a, b)
	return strconv.Atoi(combined)
}

func readInput(filename string) ([]*equation, error) {
	text, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var equations []*equation
	for _, line := range strings.Split(string(text), "\n") {
		split := strings.Split(line, ":")
		if len(split) != 2 {
			return nil, fmt.Errorf("more than 1 colon in line: %s", line)
		}
		result, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		inputs, err := shared.StrSliceToInt(strings.Split(split[1], " "))
		if err != nil {
			return nil, err
		}
		equations = append(equations, &equation{
			testValue: result,
			inputs:    inputs,
		})
	}
	return equations, nil
}
