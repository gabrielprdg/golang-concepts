package main

import (
	"fmt"
	"math"
)

// parsing types

func main() {
	var x, y int = 3, 5
	var floatValue float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(floatValue)
	fmt.Println(x, y, z, floatValue)
}
