package main

import "fmt"

func main() {
	//Declaracion de Mapas
	//var bootcampers = make(map[string]int)
	debts := map[string]int{
		"Credit":   2500,
		"Mortgage": 1500,
	}

	//Agregando un elemento al map
	debts["Pay car"] = 2500
	debts["Loans"] = 5500
	var add int
	//Recorriendo un map e imprimiendolo
	fmt.Println("Expenses    Value")
	for key, value := range debts {
		fmt.Printf("%s    %d$\n", key, value)
		add += value
	}
	fmt.Printf("Total       %d$\n", add)
	fmt.Printf("Data type position debts[\"Pay car\"]: %T \n", debts["Pay car"])
	fmt.Printf("Map data type: %T \n", debts)
}
