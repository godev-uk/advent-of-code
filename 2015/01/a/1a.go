package main

import (
	"fmt"
	"os"
	"strings"
)

func finalFloor(startFloor int, input string) int {
	currentFloor := startFloor
	instructions := strings.Split(input, "")

	for i := range instructions {
		switch instructions[i] {
		case "(":
			currentFloor++
		case ")":
			currentFloor--
		}
	}

	return currentFloor
}

func main() {
	inputBytes, _ := os.ReadFile("../2015-01-input.txt")
	inputString := string(inputBytes)

	fmt.Println(finalFloor(0, inputString))
}
