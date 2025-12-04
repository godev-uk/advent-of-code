package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	DIRECTION_LEFT  = -1
	DIRECTION_RIGHT = 1
)

const highestPosition = 99
const startPosition = 50

type Dial struct {
	direction int
	rotation  int
}

func rotate(position int, direction int, rotation int) int {
	positionCount := highestPosition + 1
	initialMod := (position + (direction * rotation)) % positionCount
	return (initialMod + positionCount) % positionCount
}

func rotateDials(position int, dials []Dial) int {
	for d := range dials {
		position = rotate(position, dials[d].direction, dials[d].rotation)
	}

	return position
}

func rotateDialsWithTarget(target int, position int, dials []Dial) int {
	targetCount := 0

	for d := range dials {
		position = rotate(position, dials[d].direction, dials[d].rotation)

		if position == target {
			targetCount++
		}
	}

	return targetCount
}

func getDial(input string) (Dial, error) {
	dial := Dial{}
	input = strings.TrimSpace(input)

	dialRegex := regexp.MustCompile(`(L|R){1}([0-9]+)`)
	if dialRegex.MatchString(input) {
		matches := dialRegex.FindStringSubmatch(input)

		if matches[1] == "L" {
			dial.direction = DIRECTION_LEFT
		} else {
			dial.direction = DIRECTION_RIGHT
		}

		dial.rotation, _ = strconv.Atoi(matches[2])

		return dial, nil
	}

	return dial, errors.New("invalid dial input")
}

func getDials(input string) []Dial {
	dials := []Dial{}

	lines := strings.Split(input, "\n")

	for lineIndex := range lines {
		dial, err := getDial(lines[lineIndex])

		if err == nil {
			dials = append(dials, dial)
		}
	}

	return dials
}

func main() {
	inputBytes, _ := os.ReadFile("../2025-01-input.txt")
	inputString := string(inputBytes)

	dials := getDials(inputString)

	fmt.Println(rotateDialsWithTarget(0, startPosition, dials))
}
