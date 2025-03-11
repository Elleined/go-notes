package main

import "fmt"

func main() {
	fmt.Println("=============================================")
	fmt.Println("LONG HAND AND SHORT HAND VARIABLE DECLARATION")
	fmt.Println("=============================================")
	longHandAndShortHandVariableDeclaration()
	fmt.Println("================== END LONG HAND AND SHORT HAND VARIABLE DECLARATION ==================")

	fmt.Println("\n\n==================")
	fmt.Println("NUMERIC DATA TYPES")
	fmt.Println("==================")
	numericDataType()
	fmt.Println("================== END NUMERIC DATA TYPE ==================")

	fmt.Println("\n\n===============")
	fmt.Println("TEXT DATA TYPES")
	fmt.Println("===============")
	textDataType()
	fmt.Println("================== END TEXT DATA TYPE ==================")

	fmt.Println("\n\n=========================")
	fmt.Println("DATA STRUCTURE DATA TYPES")
	fmt.Println("=========================")
	dataStructureDataType()
	fmt.Println("================== END DATA STRUCTURE DATA TYPE ==================")

	fmt.Println("\n\n==================")
	fmt.Println("SPECIAL DATA TYPES")
	fmt.Println("==================")
	specialDataType()
	fmt.Println("================== END OF SPECIAL DATA TYPE ==================")
}

func longHandAndShortHandVariableDeclaration() {
	// Declaring long hand syntax
	var longHand string = "Denielle"
	fmt.Printf("Printed using long hand declaration: %s \n", longHand)

	// Declaring short hand syntax
	shortHand := "Denielle"
	fmt.Printf("Printed using short hand declaration: %s \n", shortHand)
}

func numericDataType() {

}

func textDataType() {
}

func dataStructureDataType() {

}

func specialDataType() {

}
