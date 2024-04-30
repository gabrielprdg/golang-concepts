package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// we can initialize the array like this
	primes := [6]int{2, 3, 4, 6, 8, 2}
	fmt.Println(primes)
}
