package main

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
)

func TestGetGridLine(t *testing.T) {
	gridLine := []string{}
	input := ""

	assert.Equal(t, gridLine, getGridLine(input))

	gridLine = []string{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"}
	input = "MMMSXXMASM"
	assert.Equal(t, gridLine, getGridLine(input))
}

func TestGetGrid(t *testing.T) {
	grid := [][]string{}
	assert.Equal(t, grid, getGrid(""))

	input := heredoc.Doc(`
		MMMSXXMASM
		MSAMXMSMSA
		AMXSXMAAMM
		MSAMASMSMX
		XMASAMXAMM
		XXAMMXXAMA
		SMSMSASXSS
		SAXAMASAAA
		MAMMMXMMMM
		MXMXAXMASX
	`)

	grid = [][]string{
		{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
		{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
		{"A", "M", "X", "S", "X", "M", "A", "A", "M", "M"},
		{"M", "S", "A", "M", "A", "S", "M", "S", "M", "X"},
		{"X", "M", "A", "S", "A", "M", "X", "A", "M", "M"},
		{"X", "X", "A", "M", "M", "X", "X", "A", "M", "A"},
		{"S", "M", "S", "M", "S", "A", "S", "X", "S", "S"},
		{"S", "A", "X", "A", "M", "A", "S", "A", "A", "A"},
		{"M", "A", "M", "M", "M", "X", "M", "M", "M", "M"},
		{"M", "X", "M", "X", "A", "X", "M", "A", "S", "X"},
	}

	assert.Equal(t, grid, getGrid(input))
}

func TestWorldSearchCount(t *testing.T) {
	grid := [][]string{}
	search := ""
	assert.Equal(t, 0, wordSearchCount(search, grid))

	input := heredoc.Doc(`
		MMMSXXMASM
		MSAMXMSMSA
		AMXSXMAAMM
		MSAMASMSMX
		XMASAMXAMM
		XXAMMXXAMA
		SMSMSASXSS
		SAXAMASAAA
		MAMMMXMMMM
		MXMXAXMASX
	`)
	grid = getGrid(input)
	search = "XMAS"
	assert.Equal(t, 18, wordSearchCount(search, grid))
}
