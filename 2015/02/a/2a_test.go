package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPresent(t *testing.T) {
	present, err := getPresent("2x3x4")
	assert.Nil(t, err)
	assert.Equal(t, Present{
		length:        2,
		width:         3,
		height:        4,
		sideAreas:     []int{6, 8, 12},
		wrappingPaper: 58,
	}, present)

	present, err = getPresent("1x1x10")
	assert.Nil(t, err)
	assert.Equal(t, Present{
		length:        1,
		width:         1,
		height:        10,
		sideAreas:     []int{1, 10, 10},
		wrappingPaper: 43,
	}, present)

	present, err = getPresent("")
	assert.Error(t, err)
	assert.Equal(t, Present{}, present)
}
