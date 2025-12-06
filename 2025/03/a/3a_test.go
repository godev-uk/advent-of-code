package main

import (
	"os"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

var testInput string

func TestMain(m *testing.M) {
	testInputBytes, _ := os.ReadFile("../testdata/2025-03-test.txt")
	testInput = string(testInputBytes)

	os.Exit(m.Run())
}

func TestGetBatteryBank(t *testing.T) {
	batteryBank, err := getBatteryBank("")
	assert.Error(t, err)
	assert.Equal(t, BatteryBank{}, batteryBank)

	batteryBank, err = getBatteryBank("987654321111111")
	assert.Nil(t, err)
	assert.Equal(t, BatteryBank{
		batteries: []Battery{
			{
				joltage: 9,
			},
			{
				joltage: 8,
			},
			{
				joltage: 7,
			},
			{
				joltage: 6,
			},
			{
				joltage: 5,
			},
			{
				joltage: 4,
			},
			{
				joltage: 3,
			},
			{
				joltage: 2,
			},
			{
				joltage: 1,
			},
			{
				joltage: 1,
			},
			{
				joltage: 1,
			},
			{
				joltage: 1,
			},
			{
				joltage: 1,
			},
			{
				joltage: 1,
			},
			{
				joltage: 1,
			},
		},
	}, batteryBank)
}

func TestGetBatteryBanks(t *testing.T) {
	batteryBanks := getBatteryBanks(testInput)
	assert.Len(t, batteryBanks, 4)
}

func TestMaximumJoltage(t *testing.T) {
	batteryBank, _ := getBatteryBank("987654321111111")
	assert.Equal(t, 98, maximumJoltage(batteryBank))

	batteryBank, _ = getBatteryBank("811111111111119")
	assert.Equal(t, 89, maximumJoltage(batteryBank))

	batteryBank, _ = getBatteryBank("234234234234278")
	assert.Equal(t, 78, maximumJoltage(batteryBank))

	batteryBank, _ = getBatteryBank("818181911112111")
	assert.Equal(t, 92, maximumJoltage(batteryBank))
}

func TestTotalOutputJoltage(t *testing.T) {
	batteryBanks := getBatteryBanks(testInput)
	totalOutputJoltage := lo.SumBy(batteryBanks, func(batteryBank BatteryBank) int {
		return maximumJoltage(batteryBank)
	})

	assert.Equal(t, 357, totalOutputJoltage)
}
