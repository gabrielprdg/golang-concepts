package main

import "fmt"

func main() {
	// slice literals
	sliceLiterals := []int{3, 4, 5, 7, 3, 5, 20}
	fmt.Println(sliceLiterals)

	anotherSliceLiterals := []bool{false, true, false, false}
	fmt.Println(anotherSliceLiterals)

	//so we can create slice or array of type struct
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{10, false},
		{54, true},
	}

	fmt.Println(s)
}
