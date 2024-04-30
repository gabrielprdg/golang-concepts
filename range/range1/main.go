package main

import "fmt"

var pow = []int{1, 2, 4, 8, 10, 12}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
