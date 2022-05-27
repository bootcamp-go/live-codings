package main

import (
	"fmt"
)

func main() {
	var temperatura int = 20
	var humedad int = 73
	var presion float64 = 1016.9

	fmt.Printf("La temperatura es es de %d C, con una humedad del %d%, y la presion es de %f hPa", temperatura, humedad, presion)
}
