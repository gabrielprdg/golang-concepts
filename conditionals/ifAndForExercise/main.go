package main

import "fmt"

// how the square root works behind the scene

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i <= 10; i++ {
		newz := z
		z -= (z*z - x) / (2 * z)
		if z == newz {
			return z
		}
		fmt.Println(z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(10))
}
