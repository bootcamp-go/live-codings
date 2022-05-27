package main

import (
	"fmt"
)

func main() {
	fmt.Printf("The salary is: %d \n", calculateSalary(200000))
}

//1
func calculateSalary(salary int) int {

	var discount int

	if salary > 50000 && salary <= 150000 {
		discount = (salary * 17) / 100
	} else if salary > 150000 {
		discount = (salary * 27) / 100
	}
	return discount
}
