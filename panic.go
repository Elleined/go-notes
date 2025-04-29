package main

import "fmt"

func main() {
	panicFunc()
	fmt.Println("This will never be printed. Because recover is not used after a panic")

	// output
	// This will be executed first
	// This is a deferred line. Should be at the last output
	// panic: Something went wrong!
}

func panicFunc() {
	defer fmt.Println("This is a deferred line. Should be printed second")

	fmt.Println("This will be executed first")

	panic("Something went wrong!")
}
