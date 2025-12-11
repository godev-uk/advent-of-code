package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type IngredientRange struct {
	start int
	end   int
}

type Ingredient struct {
	id    int
	fresh bool
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

func getIngredient(input string) (Ingredient, error) {
	ingredient := Ingredient{}
	input = strings.TrimSpace(input)
	id, err := strconv.Atoi(input)

	if err == nil {
		ingredient.id = id
		return ingredient, nil
	}

	return ingredient, errors.New("invalid ingredient")
}

func getIngredients(input string) []Ingredient {
	ingredients := []Ingredient{}

	lines := strings.Split(input, "\n")

	for lineIndex := range lines {
		ingredient, err := getIngredient(lines[lineIndex])

		if err == nil {
			ingredients = append(ingredients, ingredient)
		}
	}

	return ingredients
}

func isFresh(id int, freshRanges []IngredientRange) bool {
	for fr := range freshRanges {
		if id >= freshRanges[fr].start && id <= freshRanges[fr].end {
			return true
		}
	}

	return false
}

func countFreshIngredients(ingredients []Ingredient, freshRanges []IngredientRange) int {
	return lo.SumBy(ingredients, func(ingredient Ingredient) int {
		if isFresh(ingredient.id, freshRanges) {
			return 1
		} else {
			return 0
		}
	})
}

func main() {
	inputBytes, _ := os.ReadFile("../2025-05-input.txt")
	inputString := string(inputBytes)

	parts := strings.Split(inputString, "\n\n")
	freshRanges := getIngredientRanges(parts[0])
	ingredients := getIngredients(parts[1])

	fmt.Println(countFreshIngredients(ingredients, freshRanges))
}
