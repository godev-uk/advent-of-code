package main

import (
	"os"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
)

var testInput string

func TestMain(m *testing.M) {
	testInputBytes, _ := os.ReadFile("../testdata/2025-04-test.txt")
	testInput = string(testInputBytes)

	os.Exit(m.Run())
}

func TestSquareDeltas(t *testing.T) {
	assert.Len(t, squareDeltas, 8)
}

func TestGetSquare(t *testing.T) {
	assert.Equal(t, Square{
		contents: UNKNOWN,
	}, getSquare(""))

	assert.Equal(t, Square{
		contents: EMPTY,
	}, getSquare("."))

	assert.Equal(t, Square{
		contents: PAPER_ROLL,
	}, getSquare("@"))
}

func TestGetLine(t *testing.T) {
	assert.Equal(t, []Square{
		{
			contents: EMPTY,
		},
		{
			contents: EMPTY,
		},
		{
			contents: PAPER_ROLL,
		},
		{
			contents: PAPER_ROLL,
		},
		{
			contents: EMPTY,
		},
		{
			contents: PAPER_ROLL,
		},
		{
			contents: PAPER_ROLL,
		},
		{
			contents: PAPER_ROLL,
		},
		{
			contents: PAPER_ROLL,
		},
		{
			contents: EMPTY,
		},
	}, getLine("..@@.@@@@."))
}

func TestGetGrid(t *testing.T) {
	gridTwoByTwo :=
		heredoc.Doc(`@.
.@`)
	assert.Equal(t, [][]Square{
		{
			{
				contents: PAPER_ROLL,
			},
			{
				contents: EMPTY,
			},
		},
		{
			{
				contents: EMPTY,
			},
			{
				contents: PAPER_ROLL,
			},
		},
	}, getGrid(gridTwoByTwo))

}

func TestCountMoveablePaperRolls(t *testing.T) {
	assert.Equal(t, 13, countMoveablePaperRolls(getGrid(testInput), 3))
}
