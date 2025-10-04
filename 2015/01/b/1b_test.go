package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTargetFloorPosition(t *testing.T) {
	assert.Equal(t, 1, targetFloorPosition(0, -1, ")"))
	assert.Equal(t, 5, targetFloorPosition(0, -1, "()())"))
	assert.Equal(t, -1, targetFloorPosition(0, -1, "((((("))
}
