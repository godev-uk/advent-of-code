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

func TestRotate(t *testing.T) {
	assert.Equal(t, 19, rotate(11, DIRECTION_RIGHT, 8))
	assert.Equal(t, 0, rotate(19, DIRECTION_LEFT, 19))
	assert.Equal(t, 99, rotate(0, DIRECTION_LEFT, 1))
	assert.Equal(t, 0, rotate(99, DIRECTION_RIGHT, 1))
}

func TestRotateDials(t *testing.T) {
	dials := getDials(testInput)

	assert.Equal(t, 32, rotateDials(50, dials))
}

func TestRotateDialsWithTarget(t *testing.T) {
	dials := getDials(testInput)

	assert.Equal(t, 3, rotateDialsWithTarget(0, 50, dials))
}

func TestGetDial(t *testing.T) {
	dial, err := getDial("L68")
	assert.Nil(t, err)
	assert.Equal(t, Dial{
		direction: DIRECTION_LEFT,
		rotation:  68,
	}, dial)

	dial, err = getDial("L30")
	assert.Nil(t, err)
	assert.Equal(t, Dial{
		direction: DIRECTION_LEFT,
		rotation:  30,
	}, dial)

	dial, err = getDial("R48")
	assert.Nil(t, err)
	assert.Equal(t, Dial{
		direction: DIRECTION_RIGHT,
		rotation:  48,
	}, dial)

	dial, err = getDial("L5")
	assert.Nil(t, err)
	assert.Equal(t, Dial{
		direction: DIRECTION_LEFT,
		rotation:  5,
	}, dial)

	dial, err = getDial("R60")
	assert.Nil(t, err)
	assert.Equal(t, Dial{
		direction: DIRECTION_RIGHT,
		rotation:  60,
	}, dial)

	dial, err = getDial("L55")
	assert.Nil(t, err)
	assert.Equal(t, Dial{
		direction: DIRECTION_LEFT,
		rotation:  55,
	}, dial)

	dial, err = getDial("L1")
	assert.Nil(t, err)
	assert.Equal(t, Dial{
		direction: DIRECTION_LEFT,
		rotation:  1,
	}, dial)

	dial, err = getDial("L99")
	assert.Nil(t, err)
	assert.Equal(t, Dial{
		direction: DIRECTION_LEFT,
		rotation:  99,
	}, dial)

	dial, err = getDial("R14")
	assert.Nil(t, err)
	assert.Equal(t, Dial{
		direction: DIRECTION_RIGHT,
		rotation:  14,
	}, dial)

	dial, err = getDial("L82")
	assert.Nil(t, err)
	assert.Equal(t, Dial{
		direction: DIRECTION_LEFT,
		rotation:  82,
	}, dial)
}

func TestGetDials(t *testing.T) {
	dials := getDials(testInput)

	assert.Equal(t, []Dial{
		{
			direction: DIRECTION_LEFT,
			rotation:  68,
		},
		{
			direction: DIRECTION_LEFT,
			rotation:  30,
		},
		{
			direction: DIRECTION_RIGHT,
			rotation:  48,
		},
		{
			direction: DIRECTION_LEFT,
			rotation:  5,
		},
		{
			direction: DIRECTION_RIGHT,
			rotation:  60,
		},
		{
			direction: DIRECTION_LEFT,
			rotation:  55,
		},
		{
			direction: DIRECTION_LEFT,
			rotation:  1,
		},
		{
			direction: DIRECTION_LEFT,
			rotation:  99,
		},
		{
			direction: DIRECTION_RIGHT,
			rotation:  14,
		},
		{
			direction: DIRECTION_LEFT,
			rotation:  82,
		},
	}, dials)
}
