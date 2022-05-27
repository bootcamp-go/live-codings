package main

import "fmt"

func main() {
	/*Llamado a la función*/
	min, max := minmax(12, 23, 54)
	fmt.Printf("Min: %d\nMax: %d\n", min, max)
}

/*Función para encontrar el MÍNIMO
y MÁXIMO entre tres números*/
func minmax(a, b, c int) (int, int) {
	/*Variables locales*/
	var min, max int
	/*Se encuentra el MÁXIMO, luego el MÍNIMO*/
	/*Primero verifica si "a" es el MÁXIMO*/
	if a > b && a > c {
		max = a
		if b < c {
			min = b
		} else {
			min = c
		}
		/*Si "a" no fue el MÁXIMO, prueba con "b"*/
	} else if b > a && b > c {
		max = b
		if a < c {
			min = a
		} else {
			min = c
		}
		/*Si "a" y "b" no fueron los MÁXIMOS, prueba con "c" */
	} else {
		max = c
		if a < b {
			min = a
		} else {
			min = b
		}
	}
	return min, max
}
