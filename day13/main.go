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
		cost := m.solve(part)
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
	prize   shared.Point
}

func (m *machine) String() string {
	return fmt.Sprintf(`Button A: X+%d, Y+%d
Button B: X+%d, Y+%d
Prize: X=%d, Y=%d
`, int(m.buttonA.X), int(m.buttonA.Y), int(m.buttonB.X), int(m.buttonB.Y), int(m.prize.X), int(m.prize.Y))
}

func (m *machine) solve(part int) int {
	prizeX := m.prize.X
	prizeY := m.prize.Y
	if part == 2 {
		prizeX += extraOffset
		prizeY += extraOffset
	}
	aPresses := (m.buttonB.Y*prizeX - m.buttonB.X*prizeY) / (m.buttonB.Y*m.buttonA.X - m.buttonA.Y*m.buttonB.X)
	bPresses := (m.buttonA.Y*prizeX - m.buttonA.X*prizeY) / (m.buttonA.Y*m.buttonB.X - m.buttonB.Y*m.buttonA.X)

	solnX := aPresses*m.buttonA.X + bPresses*m.buttonB.X
	solnY := aPresses*m.buttonA.Y + bPresses*m.buttonB.Y

	if solnX == prizeX && solnY == prizeY {
		return aPresses*aCost + bPresses*bCost
	}
	return -1
}
