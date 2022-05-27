package main

import (
	"errors"
	"fmt"
)

func calculateArea(radius int) (int, error) {
	if radius < 0 {
		return 0, errors.New("Provide Positive Number")
	}
	return radius * radius, nil
}

func main() {
	areaValue, err := calculateArea(-1)
	if err != nil {
		fmt.Println("Error Encountered...")
		return
	}

	fmt.Println(areaValue)
}
