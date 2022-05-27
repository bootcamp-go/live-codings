package main

import (
    "errors"
    "fmt"
)


const lowSalary= 10000
var ErrSalary= errors.New("salario muy bajo")

func validateSalary(salary int) error {
    if salary <= lowSalary {
        return fmt.Errorf("validateInput: %w", ErrSalary)
    }
    return nil
}

func main() {

    err := validateSalary(10000)
    if errors.Is(err, ErrSalary) {
        fmt.Println("salario bajo")
    }
}