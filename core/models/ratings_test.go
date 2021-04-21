package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAverageEmpty(t *testing.T) {
	// TODO: add more test cases
	rs := Ratings{}

	var want float64 = 0
	got := rs.GetAverage()
	assert.Equal(t, want, got)
}

func TestGetAverage(t *testing.T) {
	// TODO: add more test cases
	rs := Ratings{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
	}

	var want float64 = 3
	got := rs.GetAverage()
	assert.Equal(t, want, got)
}
