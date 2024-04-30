package main

import f "fmt"

// Slice default
func main() {
	s := []int{2, 3, 5, 6, 10}

	// here in this examples, the slice refer to previous slice
	s = s[1:3]
	f.Println(s)

	s = s[:2]
	f.Println(s)

	s = s[1:]
	f.Println(s)
}
