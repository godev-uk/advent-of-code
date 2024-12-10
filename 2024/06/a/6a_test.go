package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput string
var testRouteInput string

func TestMain(m *testing.M) {
	testInputBytes, _ := os.ReadFile("testdata/test.txt")
	testInput = string(testInputBytes)

	testRouteInputBytes, _ := os.ReadFile("testdata/test-route.txt")
	testRouteInput = string(testRouteInputBytes)

	os.Exit(m.Run())
}

func TestRotate(t *testing.T) {
	assert.Equal(t, East, rotate(North))
	assert.Equal(t, South, rotate(East))
	assert.Equal(t, West, rotate(South))
	assert.Equal(t, North, rotate(West))
}

func TestGetGridLine(t *testing.T) {
	locations, err := getGridLine("")
	assert.Equal(t, []Location{}, locations)
	assert.Error(t, err)

	locations, err = getGridLine("....#.....")
	assert.Equal(t, []Location{
		{},
		{},
		{},
		{},
		{
			entity: Obstacle,
		},
		{},
		{},
		{},
		{},
		{},
	}, locations)
	assert.Nil(t, err)

	locations, err = getGridLine("....^....#")
	assert.Equal(t, []Location{
		{},
		{},
		{},
		{},
		{
			entity:       Guard,
			guardVisited: true,
			direction:    North,
		},
		{},
		{},
		{},
		{},
		{
			entity: Obstacle,
		},
	}, locations)
	assert.Nil(t, err)
}

func TestGetGrid(t *testing.T) {
	grid := getGrid("")
	assert.Equal(t, [][]Location{}, grid)

	grid = getGrid(testInput)
	assert.Len(t, grid, 10)
	assert.Len(t, grid[0], 10)
}

func TestGuardRoute(t *testing.T) {
	// Empty grid has no guard and therefore no guard route
	grid := getGrid("")
	err := guardRoute(grid)
	assert.Error(t, err)

	grid = getGrid(testInput)
	err = guardRoute(grid)
	assert.Nil(t, err)
	assert.Equal(t, 41, countGuardVisited(grid))
}

func TestGetGridString(t *testing.T) {
	// Empty grid should have no string and no error
	grid := getGrid("")
	str, err := getGridString(grid)
	assert.Nil(t, err)
	assert.Equal(t, "", str)

	// Full grid should match both ways
	grid = getGrid(testInput)
	str, err = getGridString(grid)
	assert.Nil(t, err)
	assert.Equal(t, testInput, str)

	// After guard routing, grid should match both ways
	grid = getGrid(testInput)
	guardRoute(grid)
	str, err = getGridString(grid)
	assert.Nil(t, err)
	assert.Equal(t, testRouteInput, str)
}
