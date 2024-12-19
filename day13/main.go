package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

type today struct{}

var (
	buttonRe = regexp.MustCompile(`Button (A|B): X\+(\d+), Y\+(\d+)`)
	prizeRe  = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)
)

const (
	aCost = 3
	bCost = 1

	extraOffset = 10000000000000
)

func (t *today) ReadInput(filename string) (any, error) {
	text, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var machines []*machine
	for _, machineString := range strings.Split(string(text), "\n\n") {
		machine := &machine{}
		for _, line := range strings.Split(machineString, "\n") {
			if matches := buttonRe.FindStringSubmatch(line); len(matches) == 4 {
				x, err := strconv.Atoi(matches[2])
				if err != nil {
					return nil, err
				}
				y, err := strconv.Atoi(matches[3])
				if err != nil {
					return nil, err
				}
				switch matches[1] {
				case "A":
					machine.buttonA = shared.Point{X: x, Y: y}
				case "B":
					machine.buttonB = shared.Point{X: x, Y: y}
				}
			} else if matches := prizeRe.FindStringSubmatch(line); len(matches) == 3 {
				x, err := strconv.Atoi(matches[1])
				if err != nil {
					return nil, err
				}
				y, err := strconv.Atoi(matches[2])
				if err != nil {
					return nil, err
				}
				machine.prize = shared.Point{X: x, Y: y}
				machine.prizePart2 = shared.Point{X: x + extraOffset, Y: y + extraOffset}
			} else {
				panic(fmt.Sprintf("not sure what to do with line: %q", line))
			}
		}
		machines = append(machines, machine)
	}
	return machines, nil
}

func (t *today) Part1(input any) (int, error) {
	return implementation(input, 1)
}

func (t *today) Part2(input any) (int, error) {
	return implementation(input, 2)
}

func implementation(input any, part int) (int, error) {
	machines, ok := input.([]*machine)
	if !ok {
		return -1, fmt.Errorf("unexpected input type %T", input)
	}
	tokens := 0
	for _, m := range machines {
		cost := m.solve()
		if cost < 0 {
			continue
		}
		tokens += cost
	}
	return tokens, nil
}

type machine struct {
	buttonA shared.Point
	buttonB shared.Point

	prize      shared.Point
	prizePart2 shared.Point
}

func (m *machine) String() string {
	return fmt.Sprintf(`Button A: X+%d, Y+%d
Button B: X+%d, Y+%d
Prize: X=%d, Y=%d
`, int(m.buttonA.X), int(m.buttonA.Y), int(m.buttonB.X), int(m.buttonB.Y), int(m.prize.X), int(m.prize.Y))
}

func (m *machine) solve() int {
	options := make([][]int, 100)
	minCost := -1
	for aPresses := range len(options) {
		// initialize options
		options[aPresses] = make([]int, 100) // not great but fine
		for bPresses := range len(options[aPresses]) {
			solvedX := aPresses*m.buttonA.X+bPresses*m.buttonB.X == m.prize.X
			solvedY := aPresses*m.buttonA.Y+bPresses*m.buttonB.Y == m.prize.Y
			if !solvedX || !solvedY {
				continue
			}
			options[aPresses][bPresses] = aPresses*aCost + bPresses*bCost
			if minCost == -1 || options[aPresses][bPresses] < minCost {
				minCost = options[aPresses][bPresses]
			}
		}
	}
	return minCost
}
