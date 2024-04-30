package main

import "fmt"

type Vertex struct {
	A int
	B int
}

func main() {
	v := Vertex{20, 30}
	p := &v
	// To access the field X of a struct when we have the struct pointer p we could write (*p).X.
	// However, that notation is cumbersome, so the language permits us instead to write just p.X,
	// without the explicit dereference.
	p.A = 1e9
	fmt.Println(p)
	fmt.Println(v)
}
