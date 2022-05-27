package main

import (
	"fmt"
	"strconv"
)

func cadenaANumero(cadena string) int {
	defer func() {
		recuperado := recover()
		if recuperado != nil {
			fmt.Println(recuperado)
		}
	}()
	numero, err := strconv.Atoi(cadena)
	if err != nil {
		panic(err)
	}
	return numero
}
func main() {
	fmt.Println("Inicia")           // Inicia
	fmt.Println(cadenaANumero("2")) // 2
	// strconv.Atoi: parsing "j": invalid syntax
	fmt.Println(cadenaANumero("j")) // 0
	fmt.Println("Fin")              // Fin
}
