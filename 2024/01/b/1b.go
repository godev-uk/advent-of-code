package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type LocationPair struct {
	left     int
	right    int
	distance int
}

func getUnsortedPairs(input string) (left []int, right []int) {
	left = []int{}
	right = []int{}
	pairs := strings.Split(input, "\n")

	for p := range pairs {
		if len(pairs[p]) > 0 {
			parts := strings.Fields(pairs[p])
			leftPart, _ := strconv.Atoi(parts[0])
			rightPart, _ := strconv.Atoi(parts[1])

			left = append(left, leftPart)
			right = append(right, rightPart)
		}
	}

	return left, right
}

func main() {
	inputBytes, _ := os.ReadFile(os.Stdin.Name())
	inputString := string(inputBytes)
	left, right := getUnsortedPairs(inputString)

	similarityScore := 0

	for index := range left {
		similarityScore += left[index] * lo.Count(right, left[index])
	}

	fmt.Println(similarityScore)
}
