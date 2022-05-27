package main

import (
    
    "fmt"
)

type SalaryError struct {
    message string
}
func (e *SalaryError) Error() string {
    return fmt.Sprintf("%s El salario ingresado no alcanza el minimo imponible", e.message)
}


func validateSalary(salary int) error {
    if salary <= 150000 {
        //Mensaje personalizado con iniciacion de estructura, en caso de querer modificar el mensaje o algun tributo del struct 
		//errorSalary :=  SalaryError{
		//	message: "El salario ingresado no alcanza el minimo imponible",
		//}
		//return &errorSalary
		return &SalaryError{}
		//este return va a mostrar la linea 12
    }
    return nil
}

func main() {
	
	salary:=100
    err := validateSalary(salary)
    if err!=nil {
		fmt.Printf("%s",err.Error())
    } else {
		fmt.Println("Debe pagar impuesto")
	}
}