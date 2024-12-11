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

func TestGetCalibration(t *testing.T) {
	calibration, err := getCalibration("")
	assert.Error(t, err)
	assert.Equal(t, Calibration{}, calibration)

	calibration, err = getCalibration("190: 10 19")
	assert.Nil(t, err)
	assert.Equal(t, Calibration{
		testValue: 190,
		operands:  []int{10, 19},
	}, calibration)

	calibration, err = getCalibration("21037: 9 7 18 13")
	assert.Nil(t, err)
	assert.Equal(t, Calibration{
		testValue: 21037,
		operands:  []int{9, 7, 18, 13},
	}, calibration)
}

func TestGetCalibrations(t *testing.T) {
	assert.Equal(t, []Calibration{}, getCalibrations(""))

	assert.Equal(t, []Calibration{
		{
			testValue: 190,
			operands:  []int{10, 19},
		},
		{
			testValue: 3267,
			operands:  []int{81, 40, 27},
		},
		{
			testValue: 83,
			operands:  []int{17, 5},
		},
		{
			testValue: 156,
			operands:  []int{15, 6},
		},
		{
			testValue: 7290,
			operands:  []int{6, 8, 6, 15},
		},
		{
			testValue: 161011,
			operands:  []int{16, 10, 13},
		},
		{
			testValue: 192,
			operands:  []int{17, 8, 14},
		},
		{
			testValue: 21037,
			operands:  []int{9, 7, 18, 13},
		},
		{
			testValue: 292,
			operands:  []int{11, 6, 16, 20},
		},
	}, getCalibrations(testInput))
}

func TestCountValidCalibrations(t *testing.T) {
	assert.Equal(t, 0, countValidCalibrations(checkValidConfigurations(getCalibrations(""))))
	assert.Equal(t, 3, countValidCalibrations(checkValidConfigurations(getCalibrations(testInput))))
}

func TestTotalCalibrationResult(t *testing.T) {
	assert.Equal(t, 0, totalCalibrationResult(checkValidConfigurations(getCalibrations(""))))
	assert.Equal(t, 3749, totalCalibrationResult(checkValidConfigurations(getCalibrations(testInput))))
}
