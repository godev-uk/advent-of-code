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

func TestGetIngredient(t *testing.T) {
	ingredient, err := getIngredient("")
	assert.Error(t, err)
	assert.Equal(t, Ingredient{}, ingredient)

	ingredient, err = getIngredient("1")
	assert.Nil(t, err)
	assert.Equal(t, Ingredient{
		id:    1,
		fresh: false,
	}, ingredient)

	ingredient, err = getIngredient("5")
	assert.Nil(t, err)
	assert.Equal(t, Ingredient{
		id:    5,
		fresh: false,
	}, ingredient)

	ingredient, err = getIngredient("8")
	assert.Nil(t, err)
	assert.Equal(t, Ingredient{
		id:    8,
		fresh: false,
	}, ingredient)

	ingredient, err = getIngredient("11")
	assert.Nil(t, err)
	assert.Equal(t, Ingredient{
		id:    11,
		fresh: false,
	}, ingredient)

	ingredient, err = getIngredient("17")
	assert.Nil(t, err)
	assert.Equal(t, Ingredient{
		id:    17,
		fresh: false,
	}, ingredient)

	ingredient, err = getIngredient("32")
	assert.Nil(t, err)
	assert.Equal(t, Ingredient{
		id:    32,
		fresh: false,
	}, ingredient)
}

func TestGetIngredients(t *testing.T) {
	assert.Equal(t, []Ingredient{
		{
			id:    1,
			fresh: false,
		},
		{
			id:    5,
			fresh: false,
		},
		{
			id:    8,
			fresh: false,
		},
		{
			id:    11,
			fresh: false,
		},
		{
			id:    17,
			fresh: false,
		},
		{
			id:    32,
			fresh: false,
		},
	}, getIngredients(heredoc.Doc(`
1
5
8
11
17
32	
	`)))
}

func TestGetIsFresh(t *testing.T) {
	ingredientRanges := getIngredientRanges(heredoc.Doc(`
3-5
10-14
16-20
12-18
	`))

	assert.False(t, isFresh(1, ingredientRanges))
	assert.True(t, isFresh(5, ingredientRanges))
	assert.False(t, isFresh(8, ingredientRanges))
	assert.True(t, isFresh(11, ingredientRanges))
	assert.True(t, isFresh(17, ingredientRanges))
	assert.False(t, isFresh(32, ingredientRanges))
}

func TestCountFreshIngredients(t *testing.T) {
	ingredients := getIngredients(heredoc.Doc(`
1
5
8
11
17
32	
	`))
	freshRanges := getIngredientRanges(heredoc.Doc(`
3-5
10-14
16-20
12-18
	`))

	assert.Equal(t, 3, countFreshIngredients(ingredients, freshRanges))
}
