package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(100))
	fmt.Printf("Now you have %g problems. \n", math.Sqrt(18))
	fmt.Printf("The value of PI is %g. \n", math.Pi)
	// to round a value, we can use the one point and the number of decimal places before the variable
	fmt.Printf("The rounded value of PI is %.3g! \n", math.Pi)
}
