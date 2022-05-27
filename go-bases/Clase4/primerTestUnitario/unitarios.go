package main

import "fmt"

func main() {
	sumTotal := Sum(2, 2)
	fmt.Println(sumTotal)
}

func Sum(num1 int, num2 int) int {
	return num1 + num2
}
