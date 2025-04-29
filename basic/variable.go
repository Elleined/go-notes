package main

import "fmt"

func main() {
	fmt.Println("=============================================")
	fmt.Println("LONG HAND AND SHORT HAND VARIABLE DECLARATION")
	fmt.Println("=============================================")
	longHandAndShortHandVariableDeclaration()
	fmt.Println("================== END LONG HAND AND SHORT HAND VARIABLE DECLARATION ==================")

	fmt.Println("\n\n=============================================")
	fmt.Println("VAR and CONST KEYWORD FOR DECLARING VARIABLE")
	fmt.Println("=============================================")
	varAndConst()
	fmt.Println("================== END VAR and CONST KEYWORD FOR DECLARING VARIABLE ==================")
}

func longHandAndShortHandVariableDeclaration() {
	fmt.Print("=== long hand === \n")
	var longHand string = "Denielle"
	fmt.Printf("long hand declaration: %s \n", longHand)

	fmt.Print("=== short hand === \n")
	shortHand := "Denielle"
	fmt.Printf("short hand declaration: %s \n", shortHand)
}

func varAndConst() {
	fmt.Print("=== var === \n")
	var name string = "Red"
	fmt.Printf("Variable created via var keyword: %s \n", name)

	fmt.Print("=== const === \n")
	const AGE uint8 = 22
	fmt.Printf("Variable created via const keyword: %d \n", AGE)
}
