package main
import "fmt"

func main() {
  fmt.Println("Anonymous function sample:")
  add := func(num1, num2 int) int {
      return num1 + num2
  }
  
  sum := add(1, 2)
  fmt.Printf("Sum is: %d \n \n", sum)
  
  fmt.Println("Receiver function sample:")
  person := Person {1, "Elleined", 21}
  fmt.Printf("Person named %s is %d years old \n \n", person.Name, person.Age())
  
  fmt.Println("Closure function sample:")
  triple := multiplier(3)
  fmt.Printf("Triple multiplier: %d \n \n", triple(5))
  
  fmt.Println("Higher order function sample:")
  product := operate(5, 5, func(num1, num2 int) int {
      return num1 * num2
  })
  fmt.Printf("Result is: %d", product)
}

// Higher order function
func operate(num1, num2 int, op func(int, int) int) int {
    return op(num1, num2)
}

// Closure function
func multiplier(factor int) func (int) int {
    return func (x int) int {
        return x * factor
    }
}
  
// Receiver function
type Person struct {
    Id    int
    Name  string
    age   int
}
  
func (p Person) Age() int {
    return p.age
}
