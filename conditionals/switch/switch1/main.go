package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s. /n", os)
	}
}
