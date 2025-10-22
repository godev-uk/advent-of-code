package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Present struct {
	length        int
	width         int
	height        int
	sideAreas     []int
	wrappingPaper int
}

func getPresent(input string) (Present, error) {
	present := Present{}
	input = strings.TrimSpace(input)
	dimensions := strings.Split(input, "x")

	if len(dimensions) == 3 {
		present.length, _ = strconv.Atoi(dimensions[0])
		present.width, _ = strconv.Atoi(dimensions[1])
		present.height, _ = strconv.Atoi(dimensions[2])

		present.sideAreas = []int{
			present.length * present.width,
			present.length * present.height,
			present.width * present.height,
		}

		present.wrappingPaper = lo.SumBy(present.sideAreas, func(item int) int {
			return 2 * item
		})
		present.wrappingPaper += lo.Min(present.sideAreas)

		return present, nil
	}

	return present, errors.New("could not parse present string")
}

func main() {
	inputBytes, _ := os.ReadFile("../2015-02-input.txt")
	inputString := string(inputBytes)

	presentStrings := strings.Split(inputString, "\n")
	presents := []Present{}

	for ps := range presentStrings {
		if len(presentStrings[ps]) > 0 {
			present, _ := getPresent(presentStrings[ps])
			presents = append(presents, present)
		}
	}

	totalWrapperPaper := lo.SumBy(presents, func(present Present) int {
		return present.wrappingPaper
	})

	fmt.Println(totalWrapperPaper)
}
