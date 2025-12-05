package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput string

func TestMain(m *testing.M) {
	testInputBytes, _ := os.ReadFile("../testdata/2025-02-test.txt")
	testInput = string(testInputBytes)

	os.Exit(m.Run())
}

func TestIsValid(t *testing.T) {
	// 11-22
	assert.Equal(t, false, isValid("11"))
	assert.Equal(t, true, isValid("12"))
	assert.Equal(t, true, isValid("13"))
	assert.Equal(t, true, isValid("14"))
	assert.Equal(t, true, isValid("15"))
	assert.Equal(t, true, isValid("16"))
	assert.Equal(t, true, isValid("17"))
	assert.Equal(t, true, isValid("18"))
	assert.Equal(t, true, isValid("19"))
	assert.Equal(t, true, isValid("20"))
	assert.Equal(t, true, isValid("21"))
	assert.Equal(t, false, isValid("22"))

	// 95-115
	assert.Equal(t, true, isValid("95"))
	assert.Equal(t, true, isValid("96"))
	assert.Equal(t, true, isValid("97"))
	assert.Equal(t, true, isValid("98"))
	assert.Equal(t, false, isValid("99"))
	assert.Equal(t, true, isValid("100"))
	assert.Equal(t, true, isValid("101"))
	assert.Equal(t, true, isValid("102"))
	assert.Equal(t, true, isValid("103"))
	assert.Equal(t, true, isValid("104"))
	assert.Equal(t, true, isValid("105"))
	assert.Equal(t, true, isValid("106"))
	assert.Equal(t, true, isValid("107"))
	assert.Equal(t, true, isValid("108"))
	assert.Equal(t, true, isValid("109"))
	assert.Equal(t, true, isValid("110"))
	assert.Equal(t, true, isValid("111"))
	assert.Equal(t, true, isValid("112"))
	assert.Equal(t, true, isValid("113"))
	assert.Equal(t, true, isValid("114"))
	assert.Equal(t, true, isValid("115"))

	// 998-1012
	assert.Equal(t, true, isValid("998"))
	assert.Equal(t, true, isValid("999"))
	assert.Equal(t, true, isValid("1000"))
	assert.Equal(t, true, isValid("1001"))
	assert.Equal(t, true, isValid("1002"))
	assert.Equal(t, true, isValid("1003"))
	assert.Equal(t, true, isValid("1004"))
	assert.Equal(t, true, isValid("1005"))
	assert.Equal(t, true, isValid("1006"))
	assert.Equal(t, true, isValid("1007"))
	assert.Equal(t, true, isValid("1008"))
	assert.Equal(t, true, isValid("1009"))
	assert.Equal(t, false, isValid("1010"))
	assert.Equal(t, true, isValid("1011"))
	assert.Equal(t, true, isValid("1012"))

	// 1188511880-1188511890
	assert.Equal(t, true, isValid("1188511880"))
	assert.Equal(t, true, isValid("1188511881"))
	assert.Equal(t, true, isValid("1188511882"))
	assert.Equal(t, true, isValid("1188511883"))
	assert.Equal(t, true, isValid("1188511884"))
	assert.Equal(t, false, isValid("1188511885"))
	assert.Equal(t, true, isValid("1188511886"))
	assert.Equal(t, true, isValid("1188511887"))
	assert.Equal(t, true, isValid("1188511888"))
	assert.Equal(t, true, isValid("1188511889"))
	assert.Equal(t, true, isValid("1188511890"))

	// 222220-222224
	assert.Equal(t, true, isValid("222220"))
	assert.Equal(t, true, isValid("222221"))
	assert.Equal(t, false, isValid("222222"))
	assert.Equal(t, true, isValid("222223"))
	assert.Equal(t, true, isValid("222224"))

	// 1698522-1698528
	assert.Equal(t, true, isValid("1698522"))
	assert.Equal(t, true, isValid("1698523"))
	assert.Equal(t, true, isValid("1698524"))
	assert.Equal(t, true, isValid("1698525"))
	assert.Equal(t, true, isValid("1698526"))
	assert.Equal(t, true, isValid("1698527"))
	assert.Equal(t, true, isValid("1698528"))

	// 446443-446449
	assert.Equal(t, true, isValid("446443"))
	assert.Equal(t, true, isValid("446444"))
	assert.Equal(t, true, isValid("446445"))
	assert.Equal(t, false, isValid("446446"))
	assert.Equal(t, true, isValid("446447"))
	assert.Equal(t, true, isValid("446448"))

	// 38593856-38593862
	assert.Equal(t, true, isValid("38593856"))
	assert.Equal(t, true, isValid("38593857"))
	assert.Equal(t, true, isValid("38593858"))
	assert.Equal(t, false, isValid("38593859"))
	assert.Equal(t, true, isValid("38593860"))
	assert.Equal(t, true, isValid("38593861"))
	assert.Equal(t, true, isValid("38593862"))
}

func TestGetProductIdRanges(t *testing.T) {
	assert.Equal(t, []ProductIdRange{
		{
			ids: []ProductId{
				{
					id:    11,
					valid: false,
				},
				{
					id:    12,
					valid: true,
				},
				{
					id:    13,
					valid: true,
				},
				{
					id:    14,
					valid: true,
				},
				{
					id:    15,
					valid: true,
				},
				{
					id:    16,
					valid: true,
				},
				{
					id:    17,
					valid: true,
				},
				{
					id:    18,
					valid: true,
				},
				{
					id:    19,
					valid: true,
				},
				{
					id:    20,
					valid: true,
				},
				{
					id:    21,
					valid: true,
				},
				{
					id:    22,
					valid: false,
				},
			},
		},
		{
			ids: []ProductId{
				{
					id:    95,
					valid: true,
				},
				{
					id:    96,
					valid: true,
				},
				{
					id:    97,
					valid: true,
				},
				{
					id:    98,
					valid: true,
				},
				{
					id:    99,
					valid: false,
				},
				{
					id:    100,
					valid: true,
				},
				{
					id:    101,
					valid: true,
				},
				{
					id:    102,
					valid: true,
				},
				{
					id:    103,
					valid: true,
				},
				{
					id:    104,
					valid: true,
				},
				{
					id:    105,
					valid: true,
				},
				{
					id:    106,
					valid: true,
				},
				{
					id:    107,
					valid: true,
				},
				{
					id:    108,
					valid: true,
				},
				{
					id:    109,
					valid: true,
				},
				{
					id:    110,
					valid: true,
				},
				{
					id:    111,
					valid: true,
				},
				{
					id:    112,
					valid: true,
				},
				{
					id:    113,
					valid: true,
				},
				{
					id:    114,
					valid: true,
				},
				{
					id:    115,
					valid: true,
				},
			},
		},
	}, getProductIdRanges("11-22,95-115"))
}
