package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinalFloor(t *testing.T) {
	assert.Equal(t, 0, finalFloor(0, "(())"))
	assert.Equal(t, 0, finalFloor(0, "()()"))
	assert.Equal(t, 3, finalFloor(0, "((("))
	assert.Equal(t, 3, finalFloor(0, "(()(()("))
	assert.Equal(t, 3, finalFloor(0, "))((((("))
	assert.Equal(t, -1, finalFloor(0, "())"))
	assert.Equal(t, -1, finalFloor(0, "))("))
	assert.Equal(t, -3, finalFloor(0, ")))"))
	assert.Equal(t, -3, finalFloor(0, ")())())"))
}
