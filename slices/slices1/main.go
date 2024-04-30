package main

import "fmt"

func main() {
	numbers := [6]int{2, 4, 7, 3, 7, 10}

	//slice
	var a []int = numbers[1:4]
	fmt.Println(a)
}
