package main

import "fmt"

// Example of the "zero" values from each type

func main() {
	var i int
	var f float64
	var b bool
	var s string

	//%q ->  is used to format a string in a quoted string literal
	fmt.Printf("%v %v %v %q \n", i, f, b, s)
}
