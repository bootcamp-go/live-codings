package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	if Sum(2, 2) != 4 {
		t.Fatal("La suma de 2 + 2 debe ser igual a 4")
	}
}
