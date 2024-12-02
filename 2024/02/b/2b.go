package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Report struct {
	levels []int
	safe   bool
}

func getReports(input string) []Report {
	reports := []Report{}

	lines := strings.Split(input, "\n")

	for lineIndex := range lines {
		if len(lines[lineIndex]) > 0 {
			reports = append(reports, Report{
				levels: lo.Map(strings.Fields(lines[lineIndex]), func(item string, index int) int {
					val, _ := strconv.Atoi(item)
					return val
				}),
			})
		}
	}

	return reports
}

func areLevelsSafe(levels []int) bool {
	ascending := lo.Filter(levels, func(item int, index int) bool {
		// First item is always ascending
		if index == 0 {
			return true
		} else {
			return levels[index] > levels[index-1]
		}
	})

	descending := lo.Filter(levels, func(item int, index int) bool {
		// First item is always descending
		if index == 0 {
			return true
		} else {
			return levels[index] < levels[index-1]
		}
	})

	if len(ascending) == len(levels) || len(descending) == len(levels) {
		safeDeltas := lo.Filter(levels, func(item int, index int) bool {
			// First item always has a safe delta
			if index == 0 {
				return true
			} else {
				delta := levels[index] - levels[index-1]

				if delta < 0 {
					delta = -delta
				}

				return delta >= 1 && delta <= 3
			}
		})

		if len(safeDeltas) == len(levels) {
			return true
		}
	}

	return false
}

func isReportSafe(report Report) bool {
	// First, check all levels of this report
	if areLevelsSafe(report.levels) {
		return true
	}

	// Not all levels are safe, so try removing one at a time
	for levelIndex := range report.levels {
		// Create new levels slice based on report levels but not the current index
		subLevels := lo.Filter(report.levels, func(level int, index int) bool {
			return index != levelIndex
		})

		if areLevelsSafe(subLevels) {
			return true
		}
	}

	return false
}

func checkReports(reports []Report) []Report {
	checkedReports := make([]Report, len(reports))
	copy(checkedReports, reports)

	for cr := range checkedReports {
		checkedReports[cr].safe = isReportSafe(checkedReports[cr])
	}

	return checkedReports
}

func main() {
	inputBytes, _ := os.ReadFile(os.Stdin.Name())
	inputString := string(inputBytes)
	checkedReports := checkReports(getReports(inputString))
	safeReports := lo.Filter(checkedReports, func(checkedReport Report, index int) bool {
		return checkedReport.safe
	})

	fmt.Println(len(safeReports))
}
