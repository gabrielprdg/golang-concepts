package main

import "fmt"

type Car struct {
	Year  int
	Price float64
}

func main() {
	v := Car{2022, 50000}
	fmt.Println(v)
}
