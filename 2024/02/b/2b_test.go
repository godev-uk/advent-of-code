package main

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/stretchr/testify/assert"
)

func TestGetReports(t *testing.T) {
	reports := []Report{}

	reports = append(reports, Report{
		levels: []int{7, 6, 4, 2, 1},
		safe:   false,
	})

	reports = append(reports, Report{
		levels: []int{1, 2, 7, 8, 9},
		safe:   false,
	})

	reports = append(reports, Report{
		levels: []int{9, 7, 6, 2, 1},
		safe:   false,
	})

	reports = append(reports, Report{
		levels: []int{1, 3, 2, 4, 5},
		safe:   false,
	})

	reports = append(reports, Report{
		levels: []int{8, 6, 4, 4, 1},
		safe:   false,
	})

	reports = append(reports, Report{
		levels: []int{1, 3, 6, 7, 9},
		safe:   false,
	})

	assert.Equal(t, reports, getReports(heredoc.Doc(`
		7 6 4 2 1
		1 2 7 8 9
		9 7 6 2 1
		1 3 2 4 5
		8 6 4 4 1
		1 3 6 7 9
	`)))
}

func TestCheckReports(t *testing.T) {
	reports := []Report{}

	reports = append(reports, Report{
		levels: []int{7, 6, 4, 2, 1},
		safe:   true,
	})

	reports = append(reports, Report{
		levels: []int{1, 2, 7, 8, 9},
		safe:   false,
	})

	reports = append(reports, Report{
		levels: []int{9, 7, 6, 2, 1},
		safe:   false,
	})

	reports = append(reports, Report{
		levels: []int{1, 3, 2, 4, 5},
		safe:   true,
	})

	reports = append(reports, Report{
		levels: []int{8, 6, 4, 4, 1},
		safe:   true,
	})

	reports = append(reports, Report{
		levels: []int{1, 3, 6, 7, 9},
		safe:   true,
	})

	assert.Equal(t, reports, checkReports(getReports(heredoc.Doc(`
		7 6 4 2 1
		1 2 7 8 9
		9 7 6 2 1
		1 3 2 4 5
		8 6 4 4 1
		1 3 6 7 9
	`))))
}
