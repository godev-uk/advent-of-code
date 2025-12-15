package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	UNKNOWN = -1
	EMPTY   = iota
	BEAM_ENTRY
	BEAM
	SPLITTER
)

type Square struct {
	contents int
}

func getSquare(input string) Square {
	square := Square{}

	switch input {
	case ".":
		square.contents = EMPTY
	case "S":
		square.contents = BEAM_ENTRY
	case "|":
		square.contents = BEAM
	case "^":
		square.contents = SPLITTER
	default:
		square.contents = UNKNOWN
	}

	return square
}

func printSquare(square Square) string {
	output := ""

	switch square.contents {
	case EMPTY:
		output = "."
	case BEAM_ENTRY:
		output = "S"
	case BEAM:
		output = "|"
	case SPLITTER:
		output = "^"
	}

	return output
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

func printLine(line []Square) string {
	var sb strings.Builder

	for lineIndex := range line {
		sb.WriteString(printSquare(line[lineIndex]))
	}

	return sb.String()
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

func printGrid(grid [][]Square) string {
	var sb strings.Builder

	for row := range grid {
		sb.WriteString(printLine(grid[row]))
		sb.WriteString("\n")
	}

	return sb.String()
}

func moveBeam(grid [][]Square) bool {
	// Find current beam row - 0 is the default
	currentBeamRow := 0

	for row := range grid {
		for column := range grid[row] {
			if grid[row][column].contents == BEAM || grid[row][column].contents == BEAM_ENTRY {
				currentBeamRow = row
			}
		}
	}

	// If we are not on the last row of the grid, move down one
	// otherwise return false and do not modify the grid
	if currentBeamRow < len(grid)-1 {
		nextBeamRow := currentBeamRow + 1

		for column := range grid[currentBeamRow] {
			if grid[currentBeamRow][column].contents == BEAM || grid[currentBeamRow][column].contents == BEAM_ENTRY {
				switch grid[nextBeamRow][column].contents {
				case EMPTY:
					grid[nextBeamRow][column].contents = BEAM
				case SPLITTER:
					if column > 0 {
						grid[nextBeamRow][column-1].contents = BEAM
					}

					if column < len(grid[nextBeamRow])-1 {
						grid[nextBeamRow][column+1].contents = BEAM
					}
				}
			}
		}

		return true
	} else {
		return false
	}
}

func moveBeamToExit(grid [][]Square) {
	for moveBeam(grid) {
	}
}

func countBeamSplits(grid [][]Square) int {
	beamSplits := 0

	moveBeamToExit(grid)

	// Find all the splitters with a beam directly above them
	for row := 1; row < len(grid); row++ {
		for column := range grid[row] {
			if grid[row][column].contents == SPLITTER && grid[row-1][column].contents == BEAM {
				beamSplits++
			}
		}
	}

	return beamSplits
}

func main() {
	inputBytes, _ := os.ReadFile("../2025-07-input.txt")
	inputString := string(inputBytes)
	fmt.Println(countBeamSplits(getGrid(inputString)))
}
