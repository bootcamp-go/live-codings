package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDog(t *testing.T) {
	amountExpected := float64(100)
	amount := Dog(10)

	assert.Equal(t, amount, amountExpected)
}

func TestCat(t *testing.T) {
	amountExpected := float64(25)
	amount := Cat(5)

	assert.Equal(t, amount, amountExpected)
}

func TestMouse(t *testing.T) {
	amountExpected := float64(1.25)
	amount := Mouse(5)

	assert.Equal(t, amount, amountExpected)
}

func TestSpider(t *testing.T) {
	amountExpected := float64(1.2)
	amount := Spider(8)

	assert.Equal(t, amount, amountExpected)
}
