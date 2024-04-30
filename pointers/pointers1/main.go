package main

import "fmt"

func main() {
	i, j := 20, 2701

	// set p with i address memory
	p := &i
	fmt.Println(p)

	// (*variable) means that we are pointing to the value to the variable address memory
	*p = 57
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(p)
	fmt.Println(j)
}
