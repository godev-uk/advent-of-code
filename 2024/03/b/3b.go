package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/samber/lo"
)

type MultiplyInstruction struct {
	x       int
	y       int
	enabled bool
}

func getMultiplyInstructions(input string) []MultiplyInstruction {
	instructions := []MultiplyInstruction{}
	instructionsRegex := regexp.MustCompile(`(do\(\))|(mul\(([0-9]{1,3}),([0-9]{1,3})\))|(don\'t\(\))`)
	matches := instructionsRegex.FindAllStringSubmatch(input, -1)
	instructionsEnabled := true

	for m := range matches {
		// matches[m][0] is the matched string, as is matches[m][1]
		// matches[m][3] and matches[m][4] are the submatches (for mul())
		if matches[m][0] == "do()" {
			instructionsEnabled = true
		} else if matches[m][0] == "don't()" {
			instructionsEnabled = false
		} else {
			x, _ := strconv.Atoi(matches[m][3])
			y, _ := strconv.Atoi(matches[m][4])

			instructions = append(instructions, MultiplyInstruction{
				x:       x,
				y:       y,
				enabled: instructionsEnabled,
			})
		}
	}

	return instructions
}

func main() {
	inputBytes, _ := os.ReadFile(os.Stdin.Name())
	inputString := string(inputBytes)

	instructions := getMultiplyInstructions(inputString)

	productSum := lo.Sum(lo.Map(instructions, func(mi MultiplyInstruction, index int) int {
		if mi.enabled {
			return mi.x * mi.y
		} else {
			return 0
		}
	}))

	fmt.Println(productSum)
}
