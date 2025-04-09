# go-notes
Notes for Go

# Naming conventions
- Package names: lowercase all  
- File names: lowercase all  
- Variables and Constants:   
  - Local: camelCase  
  - Exported: PascalCase  
- Functions  
  - Local: camelCase  
  - Exported: PascalCase  
- Structs and Interface: PascalCase

# Data types
- Numeric types
  Both positive and negative numbers
  - int: default
  - int8: -128 to 127
  - int16: -32,768 to 32,767
  - int32: -2,147,483,648 to 2,147,483,647
  - int64: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
  ```go
  var age int = 18
  or
  age := 18
  ```
  Positive numbers
  - unint: default
  - uint8: 0 to 255
  - uint16: 0 to 65,535
  - uint32: 0 to 4,294,967,295
  - uint64: 0 to 18,446,744,073,709,551,615
  ```go
  var age uint = 18
  or
  age := 18
  ```
  Decimal numbers
  - float32: Approx. 6-7 decimal digits
  - float64: Approx. 15 decimal digits
  ```go
  var age float64  = 18.1
  or
  age := 18.1
  ```
  - complex64, complex 128: Real or Imaginarytnumbers
 
- Text types
  - string: String literally with double-qoutes ""
  ```go
  var name string = "Juan"
  or
  name := "Juan"
  ```
  - rune: Character with single-qoutes ``
  ```go
  var initial rune = 'J'
  or
  initial := 'J'
  ```
  - byte: byte
 
- bool: True or false
  ```go
  var isPresent bool = true
  or
  isPresent := true
  ```
- pointer: Stores the memory address of a variable (Will be elaborated more later)
  
- Data Structure Types
  - slice: Dynamic size array just like ArrayList in java
  ```go
  names := []string {
    "Juan",
    "Pedro"
  }
  ```
  - map: key-value pair
  ```go
  persons := map[string] int {
    "Juan": 25,
    "Pedro": 20
  }
  ```
  - struct: same concept with java classes
  ```go
  type Person struct {
    Name string
    Age int
  }

  person := Person {
    Name: "Juan",
    Age: 25
  }
  ```
  - interface: same concept as intefaces in java
  ```go
  type Human interface {
    Walk()
  }
  ```
  
## Most use datatypes
- int (defaults 0)
- unint
- float64 (defaults 0.0)
- string (defaults "")
- rune (defaults '')
- bool (defaults false)
- struct

## When to use specific numerical dataypes
- Unles you have a good reason. It is recommended to use the default numerical datatype `int` and `uint`
- When the performance optimization are concerned use the specific numerical datatypes

# Variables
- Please note that unused variable are not allowed in go it will cause an compilation error.
- Long hand syntax
```go
var name string = "Juan";
```

- Short hand syntax
In short hand syntax the type is inferred automatically. With the use of `:=`. Which means it automatically detects the datatype based on given value, In this example it will be inferred as string
```
name := "Juan"
```

# Constants
- constant as the name itself it cannot be changed
- Also worth noting that constant should be defined at compile time
- So constant can be only created or derived from other constants.
- Its good to declare constant in all caps.
```go
const NAME string = "Red"
```

# Format Specifiers
| Specifier  | Description                        | Example Output                        |
|------------|------------------------------------|--------------------------------------|
| `%v`       | Default format (any value)         | `42`, `true`, `{Alice 30}`           |
| `%+v`      | Struct with field names            | `{Name:Alice Age:30}`                |
| `%#v`      | Go syntax representation           | `main.Person{Name:"Alice", Age:30}`  |
| `%T`       | Type of the value                  | `int`, `string`, `main.Person`       |
| `%d`       | Integer (base 10)                  | `123`                                |
| `%b`       | Binary representation              | `1111011`                            |
| `%o`       | Octal representation               | `173`                                |
| `%x` / `%X`| Hexadecimal (lowercase / uppercase)| `7b` / `7B`                          |
| `%f`       | Floating point (decimal)           | `3.141593`                           |
| `%e` / `%E`| Scientific notation                | `3.141593e+00` / `3.141593E+00`      |
| `%s`       | String                             | `Hello`                              |
| `%q`       | Double-quoted string               | `"Hello"`                            |
| `%p`       | Pointer address                    | `0xc000012080`                       |
| `%c`       | Character (ASCII value)            | `'A'` (for `65`)                     |
| `%t`       | Boolean                            | `true`, `false`                      |


# IF with short statement
```go
Syntax:
if statement; expression {
   // body
}

if score := 10; score > 10 {
   fmt.Printf("Score is greater than %d", score)
}
```
- Here the statement is only scoped within the if statement and cannot be access outside. This is useful for the variable that are only used once for if statement one liner.

# Function
- Function Signature
```go
func name(parameters...) (returns...) {

}

// Without any parameters and returns
func foo() {

}

// With parameters and returns
// Multiple data type declaration
func foo(bar, foo int) (int, int) {
    return bar, foo
}
```

## Named Returns
- Is basically a return with a name and this was used for more code readability and named return is initialize with default value for primitives and nil with objects.
```
func bar(param1, param2 int) (ret1, ret2 int) (
    
    return ret1, ret2
}
```

#### When to use named returns
- Used them unless they add clarity for the function signature.
For example: you have a function that updates a person and return that person what you will do is.
```go
func update(person Person, updateRequest PersonUpdateRequest) (person Person) {

return person
}
```
- Here the parameter is named person and the return also named person they will conflict right? It doesnt add clarity to the function. Instead just used a regular return for this approach.
```go
func update(person Person, updateRequest PersonUpdateRequest) Person {
   return person
}
```
This is much more readable.

#### Conclusion
- Used named return when it adds clarity to your method signature unless dont just use them everywhere. Best wsy to use named return is.
```go
func fullName(firstName, lastName string) (name string) {
   return name
}
```
- Here it adds clarity that the return of fullName function returns name right?.

# Pointer
```go
/*
1. x holds the value of 10.
2. And has the memory address of 10101.
*/
x := 10
// outputs 10

/*
1. ptr holds the memory address of x (10101)
2. &x just like saying get the memory address of x (10101)
*/
ptr := &x
// outputs 10101

/*
1. *ptr just like saying get the value of memory address that ptr pointer pointing to.
2. And ptr pointer points to memory address of (10101)
3. And memory address of (10101) is x and has a value of 10.
4. Which outputs to 10.
*/
*ptr
// outputs 10
```
`*` is basically saying that get the value of that memory address. Basically saying that "What is the value of memory address that ptr pointer pointing to".

`*` And when you see `*variable` think of it just like interacting to the real object itself.
  
`&` is saying that get the memory address of the variable. Basically saying that saying that "Where is the x"

Another example: 
```go
func main() {
  x := 10
  fmt.Println("Before modifying: ", x)
  // Pass the memory address of x
  modify(&x)
  fmt.Println("After modifying: ", x)
}

func modify(num *int) {
    *num = 15
}

/*
outputs
10
15
*/
```

# Go (Golang) Roadmap: Basic to Advanced Topics
---

## **1. Introduction to Go**
- Why Go?  
- Installing Go & Setting up the Environment  
- Your First Go Program (`Hello, World!`)  
- Go Toolchain (`go run`, `go build`, `go fmt`, `go mod`)  
- Basics of Go Modules  

---

## **2. Go Basics**
- Data Types (`int`, `float`, `string`, `bool`, etc.)  
- Variables & Constants (`var`, `:=`, `const`)  
- Functions (Definition, Parameters, Return Values, Multiple Returns)  
- Control Structures (`if`, `for`, `switch`, `select`)  
- Pointers (Declaration, Dereferencing, Passing by Reference)  

---

## **3. Collections**
- Arrays  
- Slices (Creation, Manipulation, Append, Copy)  
- Maps (Creation, Accessing, Updating, Deleting)  
- Structs (Definition, Nested Structs, Methods)  

---

## **4. Functions and Methods**
- Function Types & First-Class Functions  
- Closures & Anonymous Functions  
- Methods & Receiver Functions (Value vs. Pointer Receivers)  
- Interfaces (Definition, Implementation, Type Assertions, Type Switches)  

---

## **5. Error Handling**
- Error Interface  
- Creating Custom Errors  
- Handling Errors (Returning Errors, Wrapping Errors, `errors.Is()`, `errors.As()`)  
- Panic & Recover (When to Use, Best Practices)  

---

## **6. Concurrency (Go's Highlight Feature)**
- Goroutines (Creating, Synchronization)  
- Channels (Creating, Sending, Receiving, Buffered Channels, `select` Statements)  
- WaitGroups & Mutexes (Sync Package)  
- Context Package (Cancellation, Timeouts, Deadlines)  
- Worker Pools & Patterns for Concurrency  

---

## **7. File Handling & I/O**
- Reading & Writing Files  
- Handling Directories  
- Working with JSON (Encoding/Decoding)  
- Working with CSV, XML  
- Working with HTTP (Basic Networking)  

---

## **8. Packages and Modules**
- Creating & Using Custom Packages  
- Structuring Projects  
- Dependency Management (`go mod`)  
- Publishing Modules  

---

## **9. Testing**
- Writing Unit Tests (`testing` Package)  
- Benchmarking  
- Table-Driven Tests  
- Mocking & Dependency Injection  
- Test Coverage Analysis  

---

## **10. Standard Library Deep Dive**
- `net/http` (Building Web Servers)  
- `io`, `os`, `fmt`, `strings`, `strconv`, `math`, `time`  
- Logging (`log` package, Custom Logging)  
- Working with `context` package  
- Templates (`html/template`, `text/template`)  

---

## **11. Advanced Concurrency**
- Race Conditions & Detection (`go run -race`)  
- Using `sync` and `sync/atomic` Packages  
- Channels as Streams & Pipelines  
- Concurrency Patterns (Fan-In, Fan-Out, Worker Pool)  

---

## **12. Reflection and Generics (Go 1.18+)**
- Using `reflect` Package (Type Introspection, Dynamic Calls)  
- Introduction to Generics (Type Parameters)  
- Implementing Generic Functions & Types  

---

## **13. Building APIs**
- REST API with `net/http`  
- Middleware Implementation  
- Dependency Injection in Go  
- Error Handling in APIs (Custom Error Responses)  
- Graceful Shutdown of HTTP Servers  

---

## **14. Security**
- Secure Coding Practices  
- Data Validation & Sanitization  
- Authentication & Authorization (JWT, OAuth2)  
- TLS & HTTPS Configuration  
- Preventing Common Vulnerabilities (Injection, Race Conditions, etc.)  

---

## **15. Working with Databases**
- SQL Databases (`database/sql`, `gorm`, `sqlx`)  
- NoSQL Databases (Redis, MongoDB)  
- Migrations & ORMs  
- Connection Pooling & Optimization  

---

## **16. CLI Applications**
- Building CLI Tools  
- Argument Parsing (`flag` Package, `cobra`, `urfave/cli`)  
- Error Handling & User Feedback  

---

## **17. Performance Optimization**
- Profiling (`pprof`, `trace`)  
- Memory Management & Optimization  
- Efficient Use of Goroutines  
- Optimizing I/O & Networking  

---

## **18. Deployment & Maintenance**
- Compiling for Multiple Platforms  
- Dockerizing Go Applications  
- CI/CD for Go Applications  
- Monitoring & Logging in Production  
- Versioning & Release Management  

---

## **19. Advanced Topics**
- Plugins (Dynamically Loading Code)  
- Embedding Files (Using `embed` package)  
- Working with WebAssembly (Wasm)  
- Using `cgo` for C Interoperability  
- Distributed Systems with Go (gRPC, Protobuf)  

---

## **20. Best Practices & Design Patterns**
- Idiomatic Go (`Effective Go`)  
- SOLID Principles in Go  
- Design Patterns (`Singleton`, `Factory`, `Observer`, etc.)  
- Microservices Architecture  
- Code Organization & Project Structure  
