package main

import "fmt"

/*
	Inside a function, the := short assignment can be used in a place of a var declaration
	with implicit type

	Outside the function := is not available
*/

func main() {
	var i, j = 42, 4.4
	k := 6.34
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
