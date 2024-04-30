package main

import "fmt"

func main() {
	//A defer statement defers the execution of a function until the surrounding function returns.

	fmt.Println("Hello")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Done")
}
