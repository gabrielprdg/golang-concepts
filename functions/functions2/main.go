package main

import "fmt"

/*
A function can return any number of results,
if we need to return more than one value, we need to put the value types in inside a parentesis, just like below
*/

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return x, y
}

func main() {
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}
