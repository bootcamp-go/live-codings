package main

import "fmt"

func main() {
	//Declaracion corta de un Array
	// var saludo [5]string
	greeting := [5]string{"Hello", "bootcampers"}

	//Asignando un elemerno a un Array
	greeting[2] = "how"
	greeting[3] = "are you?"

	//Imprimeindo el Slice
	fmt.Println(greeting)
	fmt.Println(greeting[0], greeting[1], greeting[2], greeting[3])
	fmt.Printf("Data type position greeting[0]: %T \n", greeting[0])
	fmt.Printf("Array data type: %T \n", greeting)
}
