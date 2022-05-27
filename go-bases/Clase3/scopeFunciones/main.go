package main

import "fmt"

var var1 = "level 1"

func main() {
	var var2 = "level 2"
	{
		var var3 = "level 3"
		fmt.Println(var3)
	}
	fmt.Println(var1)
	fmt.Println(var2)

}
