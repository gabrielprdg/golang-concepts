package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Mell Labs"] = Vertex{
		23.32344, -34.28936,
	}

	fmt.Println(m["Mell Labs"])
}
