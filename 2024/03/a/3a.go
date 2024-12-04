package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

type MultiplyInstruction struct {
	x int
	y int
}

func getMultiplyInstructions(input string) []MultiplyInstruction {
	instructions := []MultiplyInstruction{}
	instructionsRegex := regexp.MustCompile(`mul\(([0-9]{1,3}),([0-9]{1,3})\)`)
	matches := instructionsRegex.FindAllStringSubmatch(input, -1)

	for m := range matches {
		// matches[m][0] is the matched string
		// matches[m][1] and matches[m][2] are the submatches
		x, _ := strconv.Atoi(matches[m][1])
		y, _ := strconv.Atoi(matches[m][2])

		instructions = append(instructions, MultiplyInstruction{
			x: x,
			y: y,
		})
	}

	return instructions
}

func main() {
	inputBytes, _ := os.ReadFile(os.Stdin.Name())
	inputString := string(inputBytes)

	instructions := getMultiplyInstructions(inputString)

	productSum := lo.Sum(lo.Map(instructions, func(mi MultiplyInstruction, index int) int {
		return mi.x * mi.y
	}))

	fmt.Println(productSum)
}
