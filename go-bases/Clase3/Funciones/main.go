package main

import "fmt"

func main() {

	progressiveSum(3, 10)

}

func progressiveSum(num int, leng int) {
	var sum int
	for i := 0; i < leng; i++ {
		sum += num
		fmt.Println(sum)
	}
}
