package main

import (
	"fmt"
)

func main() {

	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
	)

	minFunc, err := operation(minimum)
	if err != "" {
		fmt.Println(err)
	}

	averageFunc, err := operation(average)
	if err != "" {
		fmt.Println(err)
	}

	maxFunc, err := operation(maximum)
	if err != "" {
		fmt.Println(err)
	}

	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	fmt.Printf("----------------------------\n")
	fmt.Printf("|Minimun value--------> %d   |\n", minValue)
	fmt.Printf("----------------------------\n")
	fmt.Printf("|Maximun value--------> %d   |\n", maxValue)
	fmt.Printf("----------------------------\n")
	fmt.Printf("|Average value------> %d   |\n", averageValue)
	fmt.Printf("----------------------------\n")

}

//4
func calculateMinimum(num ...int) int {
	var numMinimum int = 10000
	for _, value := range num {
		if value < numMinimum {
			numMinimum = value
		}
	}
	return numMinimum
}

func calculateAverage(num ...int) int {
	var average int
	var add int
	for _, value := range num {
		add += value
	}
	average = add / len(num)
	return average
}

func calculateMaximum(num ...int) int {
	var numMaximum int
	for _, value := range num {
		if value > numMaximum {
			numMaximum = value
		}
	}
	return numMaximum
}

func operation(calculation string) (func(...int) int, string) {

	switch calculation {
	case "minimum":
		numberMinimum := calculateMinimum
		return numberMinimum, ""
	case "average":
		numberAverage := calculateAverage
		return numberAverage, ""
	case "maximum":
		numberMaximum := calculateMaximum
		return numberMaximum, ""
	default:
		return nil, "The operation don't exist"
	}
}
