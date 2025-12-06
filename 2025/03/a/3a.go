package main

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Battery struct {
	joltage int
}

type BatteryBank struct {
	batteries []Battery
}

func getBatteryBank(input string) (BatteryBank, error) {
	batteryBank := BatteryBank{}
	input = strings.TrimSpace(input)

	bankRegex := regexp.MustCompile(`^[1-9]+$`)
	if bankRegex.MatchString(input) {
		batteryDigits := strings.Split(input, "")

		for bd := range batteryDigits {
			batteryDigit, _ := strconv.Atoi(batteryDigits[bd])
			batteryBank.batteries = append(batteryBank.batteries, Battery{
				joltage: batteryDigit,
			})
		}

		return batteryBank, nil
	}

	return batteryBank, errors.New("invalid battery bank")
}

func getBatteryBanks(input string) []BatteryBank {
	batteryBanks := []BatteryBank{}

	lines := strings.Split(input, "\n")

	for lineIndex := range lines {
		batteryBank, err := getBatteryBank(lines[lineIndex])

		if err == nil {
			batteryBanks = append(batteryBanks, batteryBank)
		}
	}

	return batteryBanks
}

func maximumJoltage(batteryBank BatteryBank) int {
	maxJoltage := 0

	// Only calculate the maximum joltage if we have at least 2 batteries
	if len(batteryBank.batteries) >= 2 {
		for i := 0; i < len(batteryBank.batteries)-1; i++ {
			for j := i + 1; j < len(batteryBank.batteries); j++ {
				currentJoltage := (batteryBank.batteries[i].joltage * 10) + batteryBank.batteries[j].joltage

				if currentJoltage > maxJoltage {
					maxJoltage = currentJoltage
				}
			}
		}
	}

	return maxJoltage
}

func main() {
	inputBytes, _ := os.ReadFile("../2025-03-input.txt")
	inputString := string(inputBytes)

	batteryBanks := getBatteryBanks(inputString)

	totalOutputJoltage := lo.SumBy(batteryBanks, func(batteryBank BatteryBank) int {
		return maximumJoltage(batteryBank)
	})

	fmt.Println(totalOutputJoltage)
}
