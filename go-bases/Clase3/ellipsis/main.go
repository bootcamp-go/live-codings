package main

import "fmt"

func main() {
	sayHello()
	sayHello("Rahul")
	sayHello("Mohit", "Rahul", "Rohit", "Johny")
}

//Al usar ellipsis siempre debe ser el ultimo
//parametro de la funcion
func sayHello(names ...string) {
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}
