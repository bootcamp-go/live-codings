package main

import (
	"fmt"
)

func main() {
	var apellido string = "Gomez"
	var edad int = 35
	boolean := "false" // Es correcta pero semanticamente deberia ser boolean := false y sin el punto y coma (Go lo elimina antes de compilar)
	var sueldo string = "45857.90"
	var nombre string = "Julián"
	fmt.Println(apellido, edad, boolean, sueldo, nombre)
}
