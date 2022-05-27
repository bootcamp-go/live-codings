package main

import "fmt"

func Welcome(message string) {
	fmt.Println("Message: ", message)
}

func main() {
	go Welcome("Welcome to Bootcamp Go") //ejecuta una funcion sobre una Go routine
	fmt.Scanln()                         //Probar que pasa si sacamos este Scanln...
	fmt.Println("The message was transmitted?")
}
