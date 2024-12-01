package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type LocationPair struct {
	left     int
	right    int
	distance int
}

func distanceBetween(left int, right int) int {
	distance := left - right

	if distance < 0 {
		distance = -distance
	}

	return distance
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

func getSortedPairs(input string) (left []int, right []int) {
	left, right = getUnsortedPairs(input)

	sort.Slice(left, func(a int, b int) bool {
		return left[a] < left[b]
	})

	sort.Slice(right, func(a int, b int) bool {
		return right[a] < right[b]
	})

	return left, right
}

func getLocationPairs(left []int, right []int) []LocationPair {
	pairs := []LocationPair{}

	for index := range left {
		pair := LocationPair{}
		pair.left = left[index]
		pair.right = right[index]
		pair.distance = distanceBetween(left[index], right[index])

		pairs = append(pairs, pair)
	}

	return pairs
}

func main() {
	inputBytes, _ := os.ReadFile(os.Stdin.Name())
	inputString := string(inputBytes)
	locationPairs := getLocationPairs(getSortedPairs(inputString))

	distances := lo.Map(locationPairs, func(lp LocationPair, index int) int {
		return lp.distance
	})

	fmt.Println(lo.Sum(distances))
}
