package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		// a = 0, b = 1 --> 0
		// a = 1, b = 2 --> 1
		// a = 2, b = 3 --> 2
		fmt.Println("result ======= ", result)
		return result
	}
}

func main() {
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
