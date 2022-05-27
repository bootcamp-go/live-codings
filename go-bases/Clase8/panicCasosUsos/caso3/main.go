package main

import (
	"fmt"
	"strconv"
)

func main() {
	var cadena = "j"
	numero, err := strconv.Atoi(cadena)
	if err != nil {
		panic(err)
	}
	fmt.Println(numero + 10)
}
