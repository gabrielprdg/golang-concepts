package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 8, 10}
	var e []int

	printSlice(s)

	fmt.Println(len(e), cap(e), e)
	if e == nil {
		fmt.Println("Nil !")
	}

	//slice the slice to give it a zero value
	s = s[:0]
	printSlice(s)

	// Extended length
	s = s[:4]
	printSlice(s)

	// drop its firsts two values
	s = s[2:]
	printSlice(s)

	// Slices can be created with the built-in make function;
	// this is how you create dynamically-sized arrays.

	//if i create simple slice, go complete the slice with  0 value from its type

	st := make([]int, 8)
	fmt.Println(st)

	fmt.Println("=====")
	a := make([]int, 7)
	printSlice(a)

	b := make([]int, 0, 3)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[1:3]
	printSlice(d)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
