package main

import (
    "errors"
    "fmt"
)



var ErrSalary= errors.New("error: el salario ingresado no alcanza el mínimo imponible")

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
       fmt.Printf("%v",ErrSalary)
    } else {
		fmt.Println("Debe pagar impuesto")
	}
}