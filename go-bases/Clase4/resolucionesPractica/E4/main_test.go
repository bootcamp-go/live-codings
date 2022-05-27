package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateMinimun(t *testing.T) {

	expectedValue := 2

	minValue := calculateMinimum(2, 3, 3, 4, 10, 2, 4, 5)

	assert.Equal(t, minValue, expectedValue)
}

func TestCalculateMaximun(t *testing.T) {
	expectedValue := 5

	maxValue := calculateMaximum(2, 3, 3, 4, 1, 2, 4, 5)

	assert.Equal(t, maxValue, expectedValue)
}

func TestCalculateAverage(t *testing.T) {
	expectedValue := 3

	avrgValue := calculateAverage(2, 3, 3, 4, 1, 2, 4, 5)

	assert.Equal(t, avrgValue, expectedValue)
}
