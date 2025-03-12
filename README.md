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
- int
- unint
- float64
- string
- rune
- bool
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
