// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import "fmt"

func main() {
    // Initializing a slice using make method with initial lenght of 5 with 0 values each
    numbers := make([]int, 5)
    
    // Appending a multiple values in slice at once
    numbers = append(numbers, 1, 2)
    
    // Reading a value in slice
    fmt.Print("Original numbers int array: ")
    printArray(&numbers)
    
    // Copy the contents of numbers into numbers1
    numbersCopy := make([]int, len(numbers))
    copy(numbersCopy, numbers)
    fmt.Print("Copy of numbers int array: ")
    printArray(&numbersCopy)
    
    // Slicing a slice
    // slice[ lower-bound : upper-bound ] 
    // Where lower-bound index is included and upper-bound index is excluded
    numbers = []int{1, 2, 3, 4, 5}
    fmt.Print("Original array before slicing: ")
    printArray(&numbers)
    
    fmt.Print("After slicing the number int array with [:2]: ")
    slicedNumbers := numbers[:2]
    printArray(&slicedNumbers)
}

func printArray(arr *[]int) {
    for _, v := range *arr {
        fmt.Printf("%d ", v)
    }
    fmt.Println()
}
