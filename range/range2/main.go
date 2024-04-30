package main

import "fmt"

func main() {
	//we can skip the index or value by assigning _

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**1
	}

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}
