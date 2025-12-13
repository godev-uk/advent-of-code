package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type IngredientRange struct {
	start int
	end   int
}

func getIngredientRange(input string) (IngredientRange, error) {
	ingredientRange := IngredientRange{}

	input = strings.TrimSpace(input)
	ids := strings.Split(input, "-")

	if len(ids) == 2 {
		ingredientRange.start, _ = strconv.Atoi(ids[0])
		ingredientRange.end, _ = strconv.Atoi(ids[1])

		return ingredientRange, nil
	}

	return ingredientRange, errors.New("invalid ingredient range")
}

func getIngredientRanges(input string) []IngredientRange {
	ingredientRanges := []IngredientRange{}

	lines := strings.Split(input, "\n")

	for lineIndex := range lines {
		ingredientRange, err := getIngredientRange(lines[lineIndex])

		if err == nil {
			ingredientRanges = append(ingredientRanges, ingredientRange)
		}
	}

	return ingredientRanges
}

func overlap(a IngredientRange, b IngredientRange) bool {
	overlapping := (b.start >= a.start && b.start <= a.end) || (b.end >= a.start && b.end <= a.end)
	enclosing := (b.start <= a.start && b.end >= a.end) || (a.start <= b.start && a.end >= b.end)
	return overlapping || enclosing
}

func merge(a IngredientRange, b IngredientRange) (IngredientRange, error) {
	irMerge := IngredientRange{}

	if overlap(a, b) {
		irMerge.start = min(a.start, b.start)
		irMerge.end = max(a.end, b.end)

		return irMerge, nil
	}

	return irMerge, errors.New("cannot merge non-overlapping ranges")
}

func mergeOne(originalRanges []IngredientRange) []IngredientRange {
	for i := 0; i < len(originalRanges)-1; i++ {
		for j := i + 1; j < len(originalRanges); j++ {
			if overlap(originalRanges[i], originalRanges[j]) {
				mergedRanges := []IngredientRange{}

				// Merge the overlapping pair
				tmpMerged, _ := merge(originalRanges[i], originalRanges[j])
				mergedRanges = append(mergedRanges, tmpMerged)

				// Add all the ranges other than the overlapping pair
				for k := range originalRanges {
					if k != i && k != j {
						mergedRanges = append(mergedRanges, originalRanges[k])
					}
				}

				return mergedRanges
			}
		}
	}

	// If we get this far, we haven't merged any ranges
	return slices.Clone(originalRanges)
}

func mergeMany(originalRanges []IngredientRange) []IngredientRange {
	mergedRanges := slices.Clone(originalRanges)

	for {
		newMergedRanges := mergeOne(mergedRanges)

		// If the number of ranges have not changed, there are no merges remaining
		if len(mergedRanges) == len(newMergedRanges) {
			return mergedRanges
		} else {
			mergedRanges = newMergedRanges
		}
	}
}

func countFreshIngredients(freshRanges []IngredientRange) int {
	return lo.SumBy(mergeMany(freshRanges), func(item IngredientRange) int {
		return (item.end - item.start) + 1
	})
}

func main() {
	inputBytes, _ := os.ReadFile("../2025-05-input.txt")
	inputString := string(inputBytes)

	parts := strings.Split(inputString, "\n\n")
	freshRanges := getIngredientRanges(parts[0])

	fmt.Println(countFreshIngredients(freshRanges))
}
