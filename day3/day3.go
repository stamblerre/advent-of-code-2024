package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type pair struct {
	a, b int
}

func main() {
	result, err := multiply("testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Result: %v\n", result)
}

func multiply(filename string) (int, error) {
	text, err := os.ReadFile(filename)
	if err != nil {
		return -1, err
	}
	r, err := regexp.Compile(`(mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`)
	if err != nil {
		return -1, err
	}
	// This regex only finds the first match, so read through all of them?
	input := string(text)
	var pairs []pair
	enabled := true
Outer:
	for {
		matches := r.FindStringSubmatch(input)
		if len(matches) == 0 {
			break Outer
		}
		match := matches[0]
		switch match {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			trimmed := strings.TrimSuffix(strings.TrimPrefix(match, "mul("), ")")
			split := strings.Split(trimmed, ",")
			if len(split) != 2 {
				return -1, fmt.Errorf("expected length 2, got %v - %v", split, len(split))
			}
			if enabled {
				a, err := strconv.Atoi(split[0])
				if err != nil {
					return -1, err
				}
				b, err := strconv.Atoi(split[1])
				if err != nil {
					return -1, err
				}
				pairs = append(pairs, pair{
					a: a,
					b: b,
				})
			}
		}
		index := strings.Index(input, match)
		if index+len(match) >= len(input) {
			break Outer
		}
		input = input[index+len(match):]
	}
	result := 0
	for _, p := range pairs {
		result += p.a * p.b
	}
	return result, nil
}
