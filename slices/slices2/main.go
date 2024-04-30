package main

import "fmt"

func main() {
	names := [5]string{
		"Rafa",
		"Gabriel",
		"Julia",
		"Iasmim",
		"Bruno",
	}

	fmt.Println(names)

	// slice dont store any data, so if i change something in the slice,
	// that change will reflect on the array as well
	a := names[0:2]
	b := names[1:3]

	fmt.Println(a, b)
	b[0] = "Bial"

	fmt.Println(a, b)
	fmt.Println(names)
}
