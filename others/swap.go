package main
import "fmt"

func main() {
  x := 10
  y := 20
  
  fmt.Println("Before swapping")
  fmt.Println("value of x: ", x)
  fmt.Println("value of y: ", y)
  
  swap(&x, &y)
  
  fmt.Println("After swapping")
  fmt.Println("value of x: ", x)
  fmt.Println("value of y: ", y)
}

func swap(x, y *int) {
    tempX := *x
    tempY := *y
    
    *y = tempX
    *x = tempY
}
