package main

import "fmt"

const Pi = 3.1416

func circule(radio float64) (area float64, perimeter float64) {
	area = Pi * radio * radio
	perimeter = 2 * Pi * radio
	return area, perimeter
}

func main() {
	a, p := circule(8)
	fmt.Println("The area of the circle is: ", a)
	fmt.Println("The perimeter of the circle is: ", p)
}
