package main

import "fmt"

// named return values
// A return statement without arguments returns the named return values

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
