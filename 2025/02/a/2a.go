package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type ProductId struct {
	id    int
	valid bool
}

type ProductIdRange struct {
	ids []ProductId
}

func isValid(id string) bool {
	valid := true

	// Only check even length IDs as odd length IDs cannot be
	// split into two repeated sequences
	if len(id)%2 == 0 {
		characters := strings.Split(id, "")
		chunks := slices.Collect(slices.Chunk(characters, len(id)/2))
		valid = !slices.Equal(chunks[0], chunks[1])
	}

	return valid
}

func getProductIdRanges(input string) []ProductIdRange {
	productIdRanges := []ProductIdRange{}

	rangeStrings := strings.Split(input, ",")

	for r := range rangeStrings {
		idStrings := strings.Split(rangeStrings[r], "-")

		productIdRange := ProductIdRange{}
		startId, _ := strconv.Atoi(idStrings[0])
		endId, _ := strconv.Atoi(idStrings[1])

		for currentId := startId; currentId <= endId; currentId++ {
			productIdRange.ids = append(productIdRange.ids, ProductId{
				id:    currentId,
				valid: isValid(strconv.Itoa(currentId)),
			})
		}

		productIdRanges = append(productIdRanges, productIdRange)
	}

	return productIdRanges
}

func main() {
	inputBytes, _ := os.ReadFile("../2025-02-input.txt")
	inputString := string(inputBytes)

	productIdRanges := getProductIdRanges(inputString)
	invalidIdSum := lo.SumBy(productIdRanges, func(productIdRange ProductIdRange) int {
		return lo.SumBy(productIdRange.ids, func(productId ProductId) int {
			if !productId.valid {
				return productId.id
			} else {
				return 0
			}
		})
	})

	fmt.Println(invalidIdSum)
}
