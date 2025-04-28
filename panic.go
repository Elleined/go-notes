package main

import "fmt"

func main() {
	defer fmt.Println("This is a deferred line. Should be printed second")

	fmt.Println("This will be executed first")

	panic("Something went wrong!")

	// output
	// This will be executed first
	// This is a deferred line. Should be at the last output
	// panic: Something went wrong!
}
