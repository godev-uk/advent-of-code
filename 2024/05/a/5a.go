package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Ordering struct {
	before int
	after  int
}

func getOrdering(input string) (Ordering, error) {
	ordering := Ordering{}
	parts := strings.Split(input, "|")

	if len(parts) == 2 {
		ordering.before, _ = strconv.Atoi(parts[0])
		ordering.after, _ = strconv.Atoi(parts[1])

		return ordering, nil
	}

	return ordering, errors.New("invalid ordering input")
}

func getOrderings(input string) []Ordering {
	orderings := []Ordering{}

	lines := strings.Split(input, "\n")

	for lineIndex := range lines {
		ordering, err := getOrdering(lines[lineIndex])

		if err == nil {
			orderings = append(orderings, ordering)
		}
	}

	return orderings
}

func getUpdate(input string) []int {
	update := []int{}

	parts := strings.Split(input, ",")

	if len(parts) >= 2 {
		for p := range parts {
			val, _ := strconv.Atoi(parts[p])
			update = append(update, val)
		}
	}

	return update
}

func getUpdates(input string) [][]int {
	updates := [][]int{}

	lines := strings.Split(input, "\n")

	for lineIndex := range lines {
		update := getUpdate(lines[lineIndex])

		if len(update) > 0 {
			updates = append(updates, update)
		}
	}

	return updates
}

func inCorrectOrder(update []int, orderings []Ordering) bool {
	expectedOrderCount := len(update) * (len(update) - 1)
	actualOrderCount := 0

	for currentIndex := range update {
		// All pages before the current one must have a 'before' match
		for beforeIndex := 0; beforeIndex < currentIndex; beforeIndex++ {
			_, ok := lo.Find(orderings, func(item Ordering) bool {
				return item.before == update[beforeIndex] && item.after == update[currentIndex]
			})

			if ok {
				actualOrderCount++
			}
		}

		// All pages after the current one must have an 'after' match
		for afterIndex := currentIndex + 1; afterIndex < len(update); afterIndex++ {
			_, ok := lo.Find(orderings, func(item Ordering) bool {
				return item.before == update[currentIndex] && item.after == update[afterIndex]
			})

			if ok {
				actualOrderCount++
			}
		}
	}

	return actualOrderCount == expectedOrderCount
}

func main() {
	inputBytes, _ := os.ReadFile(os.Stdin.Name())
	inputString := string(inputBytes)
	orderings := getOrderings(inputString)
	updates := getUpdates(inputString)

	fmt.Println(lo.Sum(lo.Map(updates, func(update []int, index int) int {
		if inCorrectOrder(update, orderings) {
			middleIndex := (len(update) - 1) / 2
			return update[middleIndex]
		} else {
			return 0
		}
	})))
}
