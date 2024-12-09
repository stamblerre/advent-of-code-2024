package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"advent-of-code-2024.com/internal/shared"
)

type today struct {
}

type equation struct {
	testValue int
	inputs    []int
}

func (t *today) ReadInput(filename string) (any, error) {
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
		inputs, err := shared.StringSliceToInt(strings.Split(split[1], " "))
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

func (t *today) Part1(input any) (int, error) {
	equations, ok := input.([]*equation)
	if !ok {
		return -1, fmt.Errorf("unexpected input format %T", equations)
	}
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

func (t *today) Part2(input any) (int, error) {
	equations, ok := input.([]*equation)
	if !ok {
		return -1, fmt.Errorf("unexpected input format %T", equations)
	}
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

func main() {
	day := &today{}
	_, _, err := shared.Run(day, "testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
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
