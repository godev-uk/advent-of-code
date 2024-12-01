package main

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
)

func TestDistanceBetween(t *testing.T) {
	assert.Equal(t, 0, distanceBetween(0, 0))
	assert.Equal(t, 1, distanceBetween(3, 4))
	assert.Equal(t, 1, distanceBetween(4, 3))
}

func TestUnsortedPairs(t *testing.T) {
	leftExpected := []int{}
	rightExpected := []int{}

	leftActual, rightActual := getUnsortedPairs("")

	assert.Equal(t, leftExpected, leftActual)
	assert.Equal(t, rightExpected, rightActual)

	leftActual, rightActual = getUnsortedPairs(heredoc.Doc(`
		3   4
		4   3
		2   5
		1   3
		3   9
		3   3
	`))

	leftExpected = []int{3, 4, 2, 1, 3, 3}
	rightExpected = []int{4, 3, 5, 3, 9, 3}

	assert.Equal(t, leftExpected, leftActual)
	assert.Equal(t, rightExpected, rightActual)
}

func TestSortedPairs(t *testing.T) {
	leftActual, rightActual := getSortedPairs(heredoc.Doc(`
		3   4
		4   3
		2   5
		1   3
		3   9
		3   3
	`))

	leftExpected := []int{1, 2, 3, 3, 3, 4}
	rightExpected := []int{3, 3, 3, 4, 5, 9}

	assert.Equal(t, leftExpected, leftActual)
	assert.Equal(t, rightExpected, rightActual)
}

func TestGetLocationPairs(t *testing.T) {
	pairs := []LocationPair{}

	pairs = append(pairs, LocationPair{
		left:     1,
		right:    3,
		distance: 2,
	})

	pairs = append(pairs, LocationPair{
		left:     2,
		right:    3,
		distance: 1,
	})

	pairs = append(pairs, LocationPair{
		left:     3,
		right:    3,
		distance: 0,
	})

	pairs = append(pairs, LocationPair{
		left:     3,
		right:    4,
		distance: 1,
	})

	pairs = append(pairs, LocationPair{
		left:     3,
		right:    5,
		distance: 2,
	})

	pairs = append(pairs, LocationPair{
		left:     4,
		right:    9,
		distance: 5,
	})

	assert.Equal(t, pairs, getLocationPairs(getSortedPairs(heredoc.Doc(`
		3   4
		4   3
		2   5
		1   3
		3   9
		3   3
	`))))
}
