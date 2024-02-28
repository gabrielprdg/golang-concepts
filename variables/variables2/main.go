package main

import "fmt"

// if an initializer is present, the type can be ommited, the variable will take the type of initializer
var g, f = 4, 5

var c, python, java bool

const d int = 23

func main() {
	// if we dont initialize a variable, the variable receive the "0" value that is different to each type of variable
	var i int
	fmt.Printf("%T \t %T \t %T \t %T \t %T \t %T \t %T \t \n", i, c, python, java, d, g, f)
}
