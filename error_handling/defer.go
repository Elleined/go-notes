package main

import "fmt"

func main() {
	deferFunc()
	// output
	// 	This will be executed first
	// This will be executed second
	// This will be executed last
}

func deferFunc() {
	defer fmt.Println("This will be executed last")
	defer fmt.Println("This will be executed second")

	fmt.Println("This will be executed first")
}
