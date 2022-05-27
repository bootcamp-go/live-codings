package main

import (
    "errors"
	"fmt"
)



var ErrSalary= errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")

func validateSalary(hours int, value int) (int,error) {
    salary:= hours*value
	if hours<80{
		return  0,ErrSalary
	}
	if salary>=150000{
		percentage:= 150000*10/100
		salary=salary-percentage
	}
	return salary,nil

}

func main() {
	

    res,err := validateSalary(10,5)
    if err!=nil {
	   fmt.Printf("%v",err)
    } else {
		fmt.Printf("%v",res)
	}
}