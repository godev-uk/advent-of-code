package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

const (
	OP_PRODUCT = iota
	OP_SUM
)

type Problem struct {
	operation int
	numbers   []int
}

func getProblems(input string) []Problem {
	lines := strings.Split(input, "\n")

	// Number of problems is equal to the number of columns of the first line
	firstLineColumns := strings.Fields(lines[0])

	problems := make([]Problem, len(firstLineColumns))

	for lineIndex := range lines {
		columns := strings.Fields(lines[lineIndex])

		for c := range columns {
			switch columns[c] {
			case "*":
				problems[c].operation = OP_PRODUCT
			case "+":
				problems[c].operation = OP_SUM
			default:
				number, _ := strconv.Atoi(columns[c])
				problems[c].numbers = append(problems[c].numbers, number)
			}
		}
	}

	return problems
}

func solve(problem Problem) int {
	switch problem.operation {
	case OP_PRODUCT:
		return lo.Product(problem.numbers)
	case OP_SUM:
		return lo.Sum(problem.numbers)
	default:
		return -1
	}
}

func grandTotal(problems []Problem) int {
	return lo.SumBy(problems, func(problem Problem) int {
		return solve(problem)
	})
}

func main() {
	inputBytes, _ := os.ReadFile("../2025-06-input.txt")
	inputString := string(inputBytes)

	fmt.Println(grandTotal(getProblems(inputString)))
}
