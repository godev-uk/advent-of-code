package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	UNKNOWN = -1
	EMPTY   = iota
	PAPER_ROLL
)

const (
	NORTH = iota
	NORTH_EAST
	EAST
	SOUTH_EAST
	SOUTH
	SOUTH_WEST
	WEST
	NORTH_WEST
)

type Square struct {
	contents int
}

type SquareDelta struct {
	direction int
	x         int
	y         int
}

var squareDeltas = []SquareDelta{
	{
		direction: NORTH,
		x:         0,
		y:         -1,
	},
	{
		direction: NORTH_EAST,
		x:         1,
		y:         -1,
	},
	{
		direction: EAST,
		x:         1,
		y:         0,
	},
	{
		direction: SOUTH_EAST,
		x:         1,
		y:         1,
	},
	{
		direction: SOUTH,
		x:         0,
		y:         1,
	},
	{
		direction: SOUTH_WEST,
		x:         -1,
		y:         1,
	},
	{
		direction: WEST,
		x:         -1,
		y:         0,
	},
	{
		direction: NORTH_WEST,
		x:         -1,
		y:         -1,
	},
}

func getSquare(input string) Square {
	square := Square{}

	switch input {
	case "@":
		square.contents = PAPER_ROLL
	case ".":
		square.contents = EMPTY
	default:
		square.contents = UNKNOWN
	}

	return square
}

func getLine(input string) []Square {
	line := []Square{}

	input = strings.Trim(input, "\n")

	squareCharacters := strings.Split(input, "")

	for sc := range squareCharacters {
		square := getSquare(squareCharacters[sc])
		if square.contents != UNKNOWN {
			line = append(line, square)
		}
	}

	return line
}

func getGrid(input string) [][]Square {
	grid := [][]Square{}

	lines := strings.Split(input, "\n")

	for lineIndex := range lines {
		if len(lines[lineIndex]) > 0 {
			grid = append(grid, getLine(lines[lineIndex]))
		}
	}

	return grid
}

func countMoveablePaperRolls(grid [][]Square, maxAdjacentRolls int) int {
	moveablePaperRolls := 0
	minX := 0
	maxX := len(grid[0]) - 1
	minY := 0
	maxY := len(grid) - 1

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			// Only check this square if it contains a paper roll
			if grid[x][y].contents == PAPER_ROLL {
				// Check for adjacent squares using deltas
				adjacentRolls := 0

				for sd := range squareDeltas {
					adjacentX := x + squareDeltas[sd].x
					adjacentY := y + squareDeltas[sd].y

					// Check that we have valid coordinates
					if adjacentX >= minX && adjacentX <= maxX && adjacentY >= minY && adjacentY <= maxY {
						if grid[adjacentX][adjacentY].contents == PAPER_ROLL {
							adjacentRolls++
						}
					}
				}

				if adjacentRolls <= maxAdjacentRolls {
					moveablePaperRolls++
				}
			}
		}
	}

	return moveablePaperRolls
}

func main() {
	inputBytes, _ := os.ReadFile("../2025-04-input.txt")
	inputString := string(inputBytes)
	grid := getGrid(inputString)
	fmt.Println(countMoveablePaperRolls(grid, 3))
}
