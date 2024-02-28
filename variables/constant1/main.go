package main

import "fmt"

const Pi float64 = 3.14

//constants cannot be declared using the := short syntax

func main() {
	const World = "Teste"
	fmt.Println("Hello", World)
	fmt.Println("Hello", Pi, "day !")

	const Truth = true
	fmt.Println("Go rules ?", Truth)
}
