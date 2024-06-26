package main

import (
	"fmt"
	"math"
)

// if i want i can declare the variable inside the if statement like below

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func sqrt(x float64) string{
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
		sqrt(-16),
	)
}
