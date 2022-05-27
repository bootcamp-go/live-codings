package main

import "fmt"

func main() {
	slice := []int{3, 4, 5, 6, 8}
	fmt.Println(slice[3]) // 6
	fmt.Println(slice[7]) // panic: runtime error: index out of range
	fmt.Println("Fin")    // no se ejecuta
}
