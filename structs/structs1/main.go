package main

import "fmt"

type Person struct {
	age  int
	name string
}

type Vertex struct {
	a int
	b int
}

func main() {
	v := Vertex{3, 4}
	v.b = 90
	fmt.Println(v)
	fmt.Println(v.a)
	fmt.Println(Person{23, "Gabriel Rodrigues"})
}
