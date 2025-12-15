package main

import (
	"os"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
)

var testInput string

func TestMain(m *testing.M) {
	testInputBytes, _ := os.ReadFile("../testdata/2025-07-test.txt")
	testInput = string(testInputBytes)

	os.Exit(m.Run())
}

func TestGetSquare(t *testing.T) {
	assert.Equal(t, Square{
		contents: UNKNOWN,
	}, getSquare(""))

	assert.Equal(t, Square{
		contents: EMPTY,
	}, getSquare("."))

	assert.Equal(t, Square{
		contents: BEAM_ENTRY,
	}, getSquare("S"))

	assert.Equal(t, Square{
		contents: BEAM,
	}, getSquare("|"))

	assert.Equal(t, Square{
		contents: SPLITTER,
	}, getSquare("^"))
}

func TestPrintSquare(t *testing.T) {
	assert.Equal(t, ".", printSquare(Square{
		contents: EMPTY,
	}))

	assert.Equal(t, "S", printSquare(Square{
		contents: BEAM_ENTRY,
	}))

	assert.Equal(t, "|", printSquare(Square{
		contents: BEAM,
	}))

	assert.Equal(t, "^", printSquare(Square{
		contents: SPLITTER,
	}))
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
			contents: BEAM_ENTRY,
		},
		{
			contents: EMPTY,
		},
		{
			contents: EMPTY,
		},
	}, getLine("..S.."))

	assert.Equal(t, []Square{
		{
			contents: EMPTY,
		},
		{
			contents: SPLITTER,
		},
		{
			contents: EMPTY,
		},
		{
			contents: SPLITTER,
		},
		{
			contents: EMPTY,
		},
	}, getLine(".^.^."))
}

func TestPrintLine(t *testing.T) {
	assert.Equal(t, "..S..", printLine([]Square{
		{
			contents: EMPTY,
		},
		{
			contents: EMPTY,
		},
		{
			contents: BEAM_ENTRY,
		},
		{
			contents: EMPTY,
		},
		{
			contents: EMPTY,
		},
	}))

	assert.Equal(t, ".^.^.", printLine([]Square{
		{
			contents: EMPTY,
		},
		{
			contents: SPLITTER,
		},
		{
			contents: EMPTY,
		},
		{
			contents: SPLITTER,
		},
		{
			contents: EMPTY,
		},
	}))
}

func TestGetGrid(t *testing.T) {
	assert.Equal(t, [][]Square{
		{
			{
				contents: EMPTY,
			},
			{
				contents: EMPTY,
			},
			{
				contents: BEAM_ENTRY,
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
				contents: EMPTY,
			},
			{
				contents: EMPTY,
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
				contents: EMPTY,
			},
			{
				contents: SPLITTER,
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
				contents: SPLITTER,
			},
			{
				contents: EMPTY,
			},
			{
				contents: EMPTY,
			},
		},
	}, getGrid(heredoc.Doc(`
..S.
....
..^.
.^..
`)))
}

func TestPrintGrid(t *testing.T) {
	assert.Equal(t, heredoc.Doc(`
..S.
....
..^.
.^..
`), printGrid([][]Square{
		{
			{
				contents: EMPTY,
			},
			{
				contents: EMPTY,
			},
			{
				contents: BEAM_ENTRY,
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
				contents: EMPTY,
			},
			{
				contents: EMPTY,
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
				contents: EMPTY,
			},
			{
				contents: SPLITTER,
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
				contents: SPLITTER,
			},
			{
				contents: EMPTY,
			},
			{
				contents: EMPTY,
			},
		},
	},
	))
}

func TestMoveBeam(t *testing.T) {
	grid := getGrid(testInput)

	assert.Equal(t, testInput, printGrid(grid))

	moveBeam(grid)
	assert.Equal(t, heredoc.Doc(`
.......S.......
.......|.......
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`), printGrid(grid))

	moveBeam(grid)
	assert.Equal(t, heredoc.Doc(`
.......S.......
.......|.......
......|^|......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`), printGrid(grid))

	moveBeam(grid)
	assert.Equal(t, heredoc.Doc(`
.......S.......
.......|.......
......|^|......
......|.|......
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`), printGrid(grid))
}

func TestMoveBeamToExit(t *testing.T) {
	grid := getGrid(testInput)
	assert.Equal(t, testInput, printGrid(grid))

	moveBeamToExit(grid)
	assert.Equal(t, heredoc.Doc(`
.......S.......
.......|.......
......|^|......
......|.|......
.....|^|^|.....
.....|.|.|.....
....|^|^|^|....
....|.|.|.|....
...|^|^|||^|...
...|.|.|||.|...
..|^|^|||^|^|..
..|.|.|||.|.|..
.|^|||^||.||^|.
.|.|||.||.||.|.
|^|^|^|^|^|||^|
|.|.|.|.|.|||.|
`), printGrid(grid))
}

func TestCountBeamSplits(t *testing.T) {
	assert.Equal(t, 21, countBeamSplits(getGrid(testInput)))
}
