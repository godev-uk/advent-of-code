package main

import (
	"fmt"
	"os"
	"strings"
)

func targetFloorPosition(startFloor int, targetFloor int, input string) int {
	currentFloor := startFloor
	instructions := strings.Split(input, "")

	for i := range instructions {
		switch instructions[i] {
		case "(":
			currentFloor++
		case ")":
			currentFloor--
		}

		if currentFloor == targetFloor {
			// range is zero indexed but positions are one indexed
			return i + 1
		}
	}

	// Never reached the target floor
	return -1
}

func main() {
	inputBytes, _ := os.ReadFile("../2015-01-input.txt")
	inputString := string(inputBytes)

	fmt.Println(targetFloorPosition(0, -1, inputString))
}
