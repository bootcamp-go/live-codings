package main

import (
	"fmt"
	"math"
)

// Aquí podemos ver la interfaz básica de una
// figura geométrica.
type geometry interface {
	area() float64
	perim() float64
}

// Para nuestro ejemplo vamos a implementar esta
// interfaz en los tipos `cuadro` y `circulo`.
type square struct {
	width, height float64
}
type circul struct {
	radio float64
}

// Para implementar una interfaz en Go, solo tenemos
// que implementar todos los metodos de la misma.
// Aquí implementamos `geometrica` en `cuadro`.
func (s square) area() float64 {
	return s.width * s.height
}
func (s square) perim() float64 {
	return 2*s.width + 2*s.height
}

// La implementación para `circulo`
func (c circul) area() float64 {
	return math.Pi * c.radio * c.radio
}
func (c circul) perim() float64 {
	return 2 * math.Pi * c.radio
}

// Si una variable es de tipo de una interfaz podemos
// llamar directamente los metodos definidos en dicha interfaz. Aquí tenemos
// un ejemplo genérico de la función `medida` que aprovecha
// esto para poder actuar sobre cualquier `geometrica`.
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	s := square{width: 3, height: 4}
	c := circul{radio: 5}

	// Los tipos `circulo` y `cuadro` implementan la
	// interfaz `geometrica` asi que podemos usar instancias de ambas
	// como parámetro de `medida`.
	measure(s)
	measure(c)
}
