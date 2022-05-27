package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax17(t *testing.T) {
	salary := 100000
	expectedtax := 17000

	tax := calculateSalary(salary)

	assert.Equal(t, tax, expectedtax)
}

func TestCalculateTax27(t *testing.T) {
	salary := 200000
	expectedtax := 54000

	tax := calculateSalary(salary)

	assert.Equal(t, tax, expectedtax)
}

func TestCalulateTax0(t *testing.T) {
	salary := 40000

	expectedtax := 0

	tax := calculateSalary(salary)

	assert.Equal(t, tax, expectedtax)
}
