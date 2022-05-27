package main

import "fmt"

func main() {
	//Declaracion corta de un Slice
	//var pair []int
	pair := []int{2, 4, 6, 8, 10}

	//Imprimiendo un rango de un Slice
	fmt.Println(pair[1:4])
	//Imprimeiendo una posicion del Slice
	fmt.Println(pair[0])

	//Agregando un elemento al Slice
	pair = append(pair, 12, 14, 16, 18, 20)
	fmt.Println(pair)

	fmt.Printf("Data type position greeting[2]: %T \n", pair[2])
	fmt.Printf("Slice data type: %T \n", pair)
}
