package main

import "fmt"

//C's while is spelled for in Go.

func main() {
	sum := 1

	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}

}
