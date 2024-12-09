package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput string

func TestMain(m *testing.M) {
	testInputBytes, _ := os.ReadFile("testdata/test.txt")
	testInput = string(testInputBytes)

	os.Exit(m.Run())
}

func TestGetOrdering(t *testing.T) {
	ordering, err := getOrdering("")
	assert.Equal(t, Ordering{}, ordering)
	assert.Error(t, err)

	ordering, err = getOrdering("47|53")
	assert.Equal(t, Ordering{
		before: 47,
		after:  53,
	}, ordering)
	assert.Nil(t, err)
}

func TestGetOrderings(t *testing.T) {
	orderings := []Ordering{}
	assert.Equal(t, orderings, getOrderings(""))
}

func TestGetUpdates(t *testing.T) {
	updates := [][]int{}

	updates = append(updates, []int{75, 47, 61, 53, 29})
	updates = append(updates, []int{97, 61, 53, 29, 13})
	updates = append(updates, []int{75, 29, 13})
	updates = append(updates, []int{75, 97, 47, 61, 53})
	updates = append(updates, []int{61, 13, 29})
	updates = append(updates, []int{97, 13, 75, 29, 47})

	assert.Equal(t, [][]int{}, getUpdates(""))
	assert.Equal(t, updates, getUpdates(testInput))
}

func TestInCorrectOrder(t *testing.T) {
	orderings := getOrderings(testInput)

	assert.True(t, inCorrectOrder(getUpdate("75,47,61,53,29"), orderings))
	assert.True(t, inCorrectOrder(getUpdate("97,61,53,29,13"), orderings))
	assert.True(t, inCorrectOrder(getUpdate("75,29,13"), orderings))
	assert.False(t, inCorrectOrder(getUpdate("75,97,47,61,53"), orderings))
	assert.False(t, inCorrectOrder(getUpdate("61,13,29"), orderings))
	assert.False(t, inCorrectOrder(getUpdate("97,13,75,29,47"), orderings))
}
