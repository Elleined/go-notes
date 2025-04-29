package main

import "fmt"

func main() {
	fmt.Println("Before happening a panic")
	deferPanicRecover()
	fmt.Println("After panic is executed and recovered! This should be printed")
}

func deferPanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v", r)
		}
	}()

	fmt.Println("This will be printed first")
	panic("Something went wrong!")
}
