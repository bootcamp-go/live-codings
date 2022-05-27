package main

import "fmt"

func main() {

	m, d := calculator(105, 7)

	fmt.Println("105 x 7 = ", m)
	fmt.Println("105 / 7 = ", d)

}

func calculator(a, b int) (mul int, div int) {

	//Asignamos simple a las variables
	mul = a * b
	div = a / b

	//Retornamos sin ningun parametro resultante
	return

}
