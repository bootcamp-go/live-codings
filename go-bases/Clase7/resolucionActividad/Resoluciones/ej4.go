package main

import (
    "fmt"
)



var ErrSalary= fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de")

func validateSalary(salary int) error {
    if salary <= 150000 {
        return  ErrSalary
    }
    return nil
}

func main() {
	
	salary:=100
    err := validateSalary(salary)
    if err!=nil {
	   fmt.Printf("%v: %v ",ErrSalary,salary)
    } else {
		fmt.Println("Debe pagar impuesto")
	}
}