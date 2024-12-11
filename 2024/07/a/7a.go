package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Calibration struct {
	testValue           int
	operands            []int
	validConfigurations int
}

func getCalibration(input string) (Calibration, error) {
	calibration := Calibration{}

	input = strings.TrimSpace(input)

	parts := strings.Split(input, ":")

	if len(parts) == 2 {
		calibration.testValue, _ = strconv.Atoi(parts[0])

		operands := strings.Fields(parts[1])

		if len(operands) >= 2 {
			for oIndex := range operands {
				operand, _ := strconv.Atoi(operands[oIndex])
				calibration.operands = append(calibration.operands, operand)
			}

			return calibration, nil
		} else {
			return calibration, errors.New("invalid operands")
		}
	}

	return calibration, errors.New("could not parse calibration string")
}

func getCalibrations(input string) []Calibration {
	calibrations := []Calibration{}

	lines := strings.Split(input, "\n")

	for lineIndex := range lines {
		calibration, err := getCalibration(lines[lineIndex])

		if err == nil {
			calibrations = append(calibrations, calibration)
		}
	}

	return calibrations
}

func checkValidConfigurations(calibrations []Calibration) []Calibration {
	for c := range calibrations {
		calculations := []int{}

		// Push first value as starting point
		calculations = append(calculations, calibrations[c].operands[0])

		for oIndex := 1; oIndex < len(calibrations[c].operands); oIndex++ {
			currentOperand := calibrations[c].operands[oIndex]
			previousCalculations := make([]int, len(calculations))
			copy(previousCalculations, calculations)

			calculations = []int{}

			for pc := range previousCalculations {
				calculations = append(calculations, previousCalculations[pc]+currentOperand)
				calculations = append(calculations, previousCalculations[pc]*currentOperand)
			}
		}

		calibrations[c].validConfigurations = lo.CountBy(calculations, func(item int) bool {
			return item == calibrations[c].testValue
		})
	}

	return calibrations
}

func countValidCalibrations(calibrations []Calibration) int {
	return lo.CountBy(calibrations, func(calibration Calibration) bool {
		return calibration.validConfigurations > 0
	})
}

func totalCalibrationResult(calibrations []Calibration) int {
	return lo.Sum(lo.FilterMap(calibrations, func(calibration Calibration, _ int) (int, bool) {
		if calibration.validConfigurations > 0 {
			return calibration.testValue, true
		} else {
			return -1, false
		}
	}))
}

func main() {
	inputBytes, _ := os.ReadFile(os.Stdin.Name())
	inputString := string(inputBytes)
	fmt.Println(totalCalibrationResult(checkValidConfigurations(getCalibrations(inputString))))
}
