package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Delta struct {
	row    int
	column int
}

func ReverseString(input string) string {
	chars := strings.Split(input, "")
	slices.Reverse(chars)
	return strings.Join(chars, "")
}

func getGridLine(input string) []string {
	gridLine := []string{}

	input = strings.TrimSpace(input)

	if len(input) > 0 {
		gridLine = strings.Split(input, "")
	}

	return gridLine
}

func getGrid(input string) [][]string {
	grid := [][]string{}

	lines := strings.Split(input, "\n")

	for lineIndex := range lines {
		if len(lines[lineIndex]) > 0 {
			grid = append(grid, getGridLine(lines[lineIndex]))
		}
	}

	return grid
}

func wordSearchCount(search string, grid [][]string) int {
	count := 0

	compassPointDeltas := make(map[string]Delta)
	compassPointDeltas["north"] = Delta{
		row:    0,
		column: -1,
	}
	compassPointDeltas["north_east"] = Delta{
		row:    1,
		column: -1,
	}
	compassPointDeltas["east"] = Delta{
		row:    1,
		column: 0,
	}
	compassPointDeltas["south_east"] = Delta{
		row:    1,
		column: 1,
	}
	compassPointDeltas["south"] = Delta{
		row:    0,
		column: 1,
	}
	compassPointDeltas["south_west"] = Delta{
		row:    -1,
		column: 1,
	}
	compassPointDeltas["west"] = Delta{
		row:    -1,
		column: 0,
	}
	compassPointDeltas["north_west"] = Delta{
		row:    -1,
		column: -1,
	}

	if len(search) > 0 {
		firstChar := search[:1]

		rowCount := len(grid)
		columnCount := 0

		if rowCount > 0 {
			columnCount = len(grid[0])
		}

		for row := range grid {
			for column := range grid[row] {
				if grid[row][column] == firstChar {
					// Look around in every direction for the search pattern
					for _, delta := range compassPointDeltas {
						str := firstChar

						for stepCount := 1; stepCount < len(search); stepCount++ {
							stepRow := row + (delta.row * stepCount)
							stepColumn := column + (delta.column * stepCount)

							if stepRow >= 0 && stepRow < rowCount && stepColumn >= 0 && stepColumn < columnCount {
								str += grid[stepRow][stepColumn]
							}
						}

						if str == search {
							// fmt.Println(str)
							count++
						}
					}
				}
			}
		}
	}

	return count
}

func main() {
	inputBytes, _ := os.ReadFile(os.Stdin.Name())
	inputString := string(inputBytes)
	grid := getGrid(inputString)

	fmt.Println(wordSearchCount("XMAS", grid))
}
