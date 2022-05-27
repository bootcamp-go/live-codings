package main

import (
	"fmt"
)

func main() {

	fmt.Printf("%v", calculateSalary("C", 180))
}

func calculateSalary(category string, minutes int) int {
	var hours int
	var salary int
	switch category {
	case "C":
		hours = minutes / 60
		salary = hours * 1000
	case "B":
		hours = minutes / 60
		salary = hours * 1500
		salary = salary + ((salary * 20) / 100)
	case "A":
		hours = minutes / 60
		salary = hours * 3000
		salary = salary + ((salary * 50) / 100)

	}
	return salary
}
