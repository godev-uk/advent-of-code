package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput string

func TestMain(m *testing.M) {
	testInputBytes, _ := os.ReadFile("../testdata/2025-06-test.txt")
	testInput = string(testInputBytes)

	os.Exit(m.Run())
}

func TestGetProblems(t *testing.T) {
	assert.Equal(t, []Problem{
		{
			operation: OP_PRODUCT,
			numbers: []int{
				123,
				45,
				6,
			},
		},
		{
			operation: OP_SUM,
			numbers: []int{
				328,
				64,
				98,
			},
		},
		{
			operation: OP_PRODUCT,
			numbers: []int{
				51,
				387,
				215,
			},
		},
		{
			operation: OP_SUM,
			numbers: []int{
				64,
				23,
				314,
			},
		},
	}, getProblems(testInput))
}

func TestSolve(t *testing.T) {
	assert.Equal(t, 33210, solve(Problem{
		operation: OP_PRODUCT,
		numbers: []int{
			123,
			45,
			6,
		},
	}))

	assert.Equal(t, 490, solve(Problem{
		operation: OP_SUM,
		numbers: []int{
			328,
			64,
			98,
		},
	}))

	assert.Equal(t, 4243455, solve(Problem{
		operation: OP_PRODUCT,
		numbers: []int{
			51,
			387,
			215,
		},
	}))

	assert.Equal(t, 401, solve(Problem{
		operation: OP_SUM,
		numbers: []int{
			64,
			23,
			314,
		},
	}))
}

func TestGrandTotal(t *testing.T) {
	assert.Equal(t, 4277556, grandTotal(getProblems(testInput)))
}
