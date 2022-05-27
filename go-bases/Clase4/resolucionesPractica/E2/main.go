package main

import (
	"fmt"
)

func main() {

	fmt.Printf("%v", calculateAverage(5, 6, 7, 10))
}

//2 calcular promedio recibe solo numero enteros, ya que uint verifica que este sea true
func calculateAverage(values ...uint) float32 {

	var add float32

	for _, value := range values {
		add += float32(value)
	}

	final := add / float32(len(values))
	return final
}
