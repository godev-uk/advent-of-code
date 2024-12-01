package main

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
)

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
