package main

import (
	"os"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
)

var testInput string

func TestMain(m *testing.M) {
	testInputBytes, _ := os.ReadFile("../testdata/2025-05-test.txt")
	testInput = string(testInputBytes)

	os.Exit(m.Run())
}

func TestGetIngredientRange(t *testing.T) {
	ingredientRange, err := getIngredientRange("")
	assert.Error(t, err)
	assert.Equal(t, IngredientRange{}, ingredientRange)

	ingredientRange, err = getIngredientRange("3-5")
	assert.Nil(t, err)
	assert.Equal(t, IngredientRange{
		start: 3,
		end:   5,
	}, ingredientRange)

	ingredientRange, err = getIngredientRange("10-14")
	assert.Nil(t, err)
	assert.Equal(t, IngredientRange{
		start: 10,
		end:   14,
	}, ingredientRange)

	ingredientRange, err = getIngredientRange("16-20")
	assert.Nil(t, err)
	assert.Equal(t, IngredientRange{
		start: 16,
		end:   20,
	}, ingredientRange)

	ingredientRange, err = getIngredientRange("12-18")
	assert.Nil(t, err)
	assert.Equal(t, IngredientRange{
		start: 12,
		end:   18,
	}, ingredientRange)
}

func TestGetIngredientRanges(t *testing.T) {
	assert.Equal(t, []IngredientRange{
		{
			start: 3,
			end:   5,
		},
		{
			start: 10,
			end:   14,
		},
		{
			start: 16,
			end:   20,
		},
		{
			start: 12,
			end:   18,
		},
	}, getIngredientRanges(heredoc.Doc(`
3-5
10-14
16-20
12-18
	`)))
}

func TestCountFreshIngredients(t *testing.T) {
	freshRanges := getIngredientRanges(heredoc.Doc(`
3-5
10-14
16-20
12-18
	`))

	assert.Equal(t, 14, countFreshIngredients(freshRanges))

	freshRanges = getIngredientRanges(heredoc.Doc(`
284509448628766-289381261597719
417436534565209-420984393851206
	`))
	assert.Equal(t, 8419672254952, countFreshIngredients(freshRanges))

	freshRanges = getIngredientRanges(heredoc.Doc(`
3-5
10-14
16-20
12-18
9-21
	`))
	assert.Equal(t, 16, countFreshIngredients(freshRanges))
}

func TestOverlap(t *testing.T) {
	assert.False(t, overlap(
		IngredientRange{
			start: 3,
			end:   5,
		},
		IngredientRange{
			start: 10,
			end:   14,
		},
	))

	assert.False(t, overlap(
		IngredientRange{
			start: 10,
			end:   14,
		},
		IngredientRange{
			start: 16,
			end:   20,
		},
	))

	assert.True(t, overlap(
		IngredientRange{
			start: 16,
			end:   20,
		},
		IngredientRange{
			start: 12,
			end:   18,
		},
	))

	assert.True(t, overlap(
		IngredientRange{
			start: 12,
			end:   18,
		},
		IngredientRange{
			start: 16,
			end:   20,
		},
	))

	assert.False(t, overlap(
		IngredientRange{
			start: 3,
			end:   5,
		},
		IngredientRange{
			start: 9,
			end:   21,
		},
	))

	assert.True(t, overlap(
		IngredientRange{
			start: 12,
			end:   20,
		},
		IngredientRange{
			start: 9,
			end:   21,
		},
	))
}

func TestMerge(t *testing.T) {
	irMerge, err := merge(
		IngredientRange{
			start: 3,
			end:   5,
		},
		IngredientRange{
			start: 10,
			end:   14,
		},
	)
	assert.Error(t, err)
	assert.Equal(t, IngredientRange{}, irMerge)

	irMerge, err = merge(
		IngredientRange{
			start: 16,
			end:   20,
		},
		IngredientRange{
			start: 12,
			end:   18,
		},
	)
	assert.Nil(t, err)
	assert.Equal(t, IngredientRange{
		start: 12,
		end:   20,
	}, irMerge)

	irMerge, err = merge(
		IngredientRange{
			start: 12,
			end:   18,
		},
		IngredientRange{
			start: 16,
			end:   20,
		},
	)
	assert.Nil(t, err)
	assert.Equal(t, IngredientRange{
		start: 12,
		end:   20,
	}, irMerge)
}

func TestMergeMany(t *testing.T) {
	assert.ElementsMatch(t, []IngredientRange{
		{
			start: 3,
			end:   5,
		},
		{
			start: 10,
			end:   20,
		},
	}, mergeMany([]IngredientRange{
		{
			start: 3,
			end:   5,
		},
		{
			start: 10,
			end:   14,
		},
		{
			start: 16,
			end:   20,
		},
		{
			start: 12,
			end:   18,
		},
	}))

	assert.ElementsMatch(t, []IngredientRange{
		{
			start: 284509448628766,
			end:   289381261597719,
		},
		{
			start: 417436534565209,
			end:   420984393851206,
		},
	}, mergeMany([]IngredientRange{
		{
			start: 284509448628766,
			end:   289381261597719,
		},
		{
			start: 417436534565209,
			end:   420984393851206,
		},
	}))

	assert.ElementsMatch(t, []IngredientRange{
		{
			start: 3,
			end:   5,
		},
		{
			start: 10,
			end:   20,
		},
	}, mergeMany([]IngredientRange{
		{
			start: 3,
			end:   5,
		},
		{
			start: 10,
			end:   14,
		},
		{
			start: 16,
			end:   20,
		},
		{
			start: 12,
			end:   18,
		},
	}))
}
