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
		x: 2,
		y: 4,
	})

	instructions = append(instructions, MultiplyInstruction{
		x: 5,
		y: 5,
	})

	instructions = append(instructions, MultiplyInstruction{
		x: 11,
		y: 8,
	})

	instructions = append(instructions, MultiplyInstruction{
		x: 8,
		y: 5,
	})

	assert.Equal(t, instructions, getMultiplyInstructions("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"))
}
