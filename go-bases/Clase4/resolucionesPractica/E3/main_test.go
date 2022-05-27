package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSalaryCategoryA(t *testing.T) {

	expectedSalary := 13500

	salary := calculateSalary("A", 180)

	assert.Equal(t, salary, expectedSalary)
}

func TestSalaryCategoryB(t *testing.T) {

	expectedSalary := 5400

	salary := calculateSalary("B", 180)

	assert.Equal(t, salary, expectedSalary)
}

func TestSalaryCategoryC(t *testing.T) {

	expectedSalary := 3000

	salary := calculateSalary("C", 180)

	assert.Equal(t, salary, expectedSalary)
}
