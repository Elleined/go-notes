package main

import "fmt"

// This struct.go shows the example of using the ff.
// 1. Receiver function
// 2. Method overriding with interface
// 3. Method overriding or shadowing with struct
// 4. Anonymous embedded struct or struct within a struct
// 5. Accessing the field of embedded struct
// 6. Accessing the method of embedded struct

type Walk interface {
	Walk()
}

type Human struct{}

func (h *Human) Speak() {
	fmt.Println("General speaking...")
}

type Person struct {
	Human     // Composition
	Id        int
	FirstName string
	LastName  string
	Address   // Aggregation
}

// Speak function comes from Human struct and trying to method overidding or method shadowing
func (p *Person) Speak() {
	fmt.Println("Person is speaking... and should not be generally speaking as declared in Human struct")
}

// Walk function comes from walk interface and implemented here
func (p *Person) Walk() {
	fmt.Printf("%s is walking... \n", p.FullName())
}

// Receiver function for person
func (p *Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

type Address struct {
	Province  string
	City      string
	Baranggay string
}

// Receiver function for address
func (a *Address) GetFullAddress() string {
	return a.Baranggay + " " + a.Province + ", " + a.Baranggay
}

func main() {
	person := Person{
		Id:        1,
		FirstName: "Denielle",
		LastName:  "De Guzman",
		Address: Address{
			Province:  "Nueva Ecija",
			City:      "Cabanatuan City",
			Baranggay: "Sumacab Este",
		},
	}
	fmt.Println("(Receiver function) Person's full name:", person.FullName())
	fmt.Print("(Method override from Walk interface): ")
	person.Walk()

	fmt.Print("(Method override from Human struct): ")
	person.Speak()

	fmt.Println("(Anonymous struct field accessing) Person's province:", person.Province)
	fmt.Println("(Anonymous struct method accessing) Person's full address:", person.GetFullAddress())
}
