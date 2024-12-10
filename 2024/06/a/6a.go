package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/samber/lo"
)

const (
	North = iota
	East
	South
	West
)

const (
	Empty = iota
	Obstacle
	Guard
)

var locationCharacters = map[string]Location{
	"#": {
		entity: Obstacle,
	},
	".": {
		entity: Empty,
	},
	"^": {
		entity:       Guard,
		direction:    North,
		guardVisited: true,
	},
	">": {
		entity:       Guard,
		direction:    East,
		guardVisited: true,
	},
	"v": {
		entity:       Guard,
		direction:    South,
		guardVisited: true,
	},
	"<": {
		entity:       Guard,
		direction:    West,
		guardVisited: true,
	},
	"X": {
		entity:       Empty,
		guardVisited: true,
	},
}

type Delta struct {
	row    int
	column int
}

var compassPointDeltas = map[int]Delta{
	North: {
		row:    -1,
		column: 0,
	},
	East: {
		row:    0,
		column: 1,
	},
	South: {
		row:    1,
		column: 0,
	},
	West: {
		row:    0,
		column: -1,
	},
}

type Location struct {
	entity       int
	direction    int
	guardVisited bool
}

func getGridLine(input string) ([]Location, error) {
	locations := []Location{}

	input = strings.TrimSpace(input)

	re := regexp.MustCompile(`^([#\.\<\>\^v]+)$`)

	if re.MatchString(input) {
		for c := range input {
			locations = append(locations, locationCharacters[string(input[c])])
		}

		return locations, nil
	}

	return locations, errors.New("invalid grid line input")
}

func getGrid(input string) [][]Location {
	grid := [][]Location{}

	lines := strings.Split(input, "\n")

	for lineIndex := range lines {
		locations, err := getGridLine(lines[lineIndex])

		if err == nil {
			grid = append(grid, locations)
		}
	}

	return grid
}

func guardRoute(grid [][]Location) error {
	guardLocation := Location{}
	guardRow := -1
	guardColumn := -1
	guardFound := false

	// First, find the guard
	for row := 0; row < len(grid) && !guardFound; row++ {
		for column := 0; column < len(grid[row]) && !guardFound; column++ {
			if grid[row][column].entity == Guard {
				guardFound = true
				guardLocation = grid[row][column]
				guardRow = row
				guardColumn = column
			}
		}
	}

	if !guardFound {
		return errors.New("no guard found")
	}

	// Move the guard until they go outside of the map
	rowMax := len(grid) - 1
	columnMax := len(grid[0]) - 1
	routeComplete := false
	for !routeComplete {
		// Try and move one step in the direction the guard is facing
		guardMovedRow := guardRow + compassPointDeltas[guardLocation.direction].row
		guardMovedColumn := guardColumn + compassPointDeltas[guardLocation.direction].column

		// Only update the grid if the guard is still on it
		if guardMovedRow >= 0 && guardMovedRow <= rowMax && guardMovedColumn >= 0 && guardMovedColumn <= columnMax {
			if grid[guardMovedRow][guardMovedColumn].entity == Obstacle {
				// Rotate the guard but keep position the same
				guardLocation.direction = rotate(guardLocation.direction)
			} else {
				// Guard can move to the new location
				grid[guardRow][guardColumn].entity = Empty
				guardRow = guardMovedRow
				guardColumn = guardMovedColumn
				grid[guardRow][guardColumn].entity = Guard
				grid[guardRow][guardColumn].guardVisited = true
			}
		} else {
			// Guard has left the grid, don't forget to mark their
			// final location as visited
			grid[guardRow][guardColumn].entity = Empty
			grid[guardRow][guardColumn].guardVisited = true
			routeComplete = true
		}
	}

	return nil
}

func countGuardVisited(grid [][]Location) int {
	count := 0

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			if grid[row][column].guardVisited {
				count++
			}
		}
	}

	return count
}

func getGridString(grid [][]Location) (string, error) {
	str := ""
	characterLocations := lo.Invert(locationCharacters)

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			c, ok := characterLocations[grid[row][column]]

			if ok {
				str += c
			} else {
				return "", errors.New("invalid entity in grid")
			}
		}

		str += "\n"
	}

	str = strings.TrimSpace(str)

	return str, nil
}

func rotate(direction int) int {
	switch direction {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	default:
		return -1
	}
}

func main() {
	inputBytes, _ := os.ReadFile(os.Stdin.Name())
	inputString := string(inputBytes)
	grid := getGrid(inputString)
	guardRoute(grid)
	fmt.Println(countGuardVisited(grid))
}
