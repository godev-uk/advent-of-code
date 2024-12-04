package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMultiplyInstructions(t *testing.T) {
	instructions := []MultiplyInstruction{}
	assert.Equal(t, instructions, getMultiplyInstructions(""))

	instructions = []MultiplyInstruction{}

	instructions = append(instructions, MultiplyInstruction{
		x:       2,
		y:       4,
		enabled: true,
	})

	instructions = append(instructions, MultiplyInstruction{
		x:       5,
		y:       5,
		enabled: false,
	})

	instructions = append(instructions, MultiplyInstruction{
		x:       11,
		y:       8,
		enabled: false,
	})

	instructions = append(instructions, MultiplyInstruction{
		x:       8,
		y:       5,
		enabled: true,
	})

	assert.Equal(t, instructions, getMultiplyInstructions("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"))
}
