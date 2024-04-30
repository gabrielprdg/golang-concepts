package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{3, 2}
	v2 = Vertex{X: 2}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
