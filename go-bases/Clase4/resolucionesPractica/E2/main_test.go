package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateAverage(t *testing.T) {
	expectedAverage := float32(7)

	average := calculateAverage(5, 6, 7, 10)

	assert.Equal(t, expectedAverage, average)
}
