/*
package main is a entry point of our application
and it specifies the name of the package
*/

package main

import "fmt"

// in a function we have to declare the parameters types and the return type

func add(x int, y int) int {
	return x + y
}

/*
When two or more consecutive named function parameters share a type,
you can omit the type from all but the last.
In this example below we shortened x int, y int, z int to x,y,z int
*/

func multiple(x, y, z int) int {
	return x * y * z
}

func sayHello(name string) {
	fmt.Printf("Hello, wellcome %s !\n", name)
}

func main() {
	fmt.Println(add(14, 63))
	sayHello("Gabriel Rodrigues")
	fmt.Println(multiple(4, 6, 33))
}
