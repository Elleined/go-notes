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
	var signedInt8 int8 = 127
	fmt.Printf("\t Signed Integer 8 with range of -128 to 127: %v \n", signedInt8)

	var signedInt16 int16 = 32767
	fmt.Printf("\t Signed Integer 16 with range of -32,768 to 32767: %v \n", signedInt16)

	var signedInt32 int32 = 2147483647
	fmt.Printf("\t Signed Integer 32 with range of -2,147,483,648 to 2147483647: %v \n", signedInt32)

	var signedInt64 int64 = 9223372036854775807
	fmt.Printf("\t Signed Integer 64 with range of -9,223,372,036,854,775,808 to 9223372036854775807: %v \n", signedInt64)
	fmt.Println("\t ====================================================")

	var unSignedInt8 uint8 = 255
	fmt.Printf("\t Unsigned Integer 8 with range of 0 to 255: %v \n", unSignedInt8)

	var unSignedInt16 uint16 = 65535
	fmt.Printf("\t Unsigned Integer 16 with range of 0 to 65535: %v \n", unSignedInt16)

	var unSignedInt32 uint32 = 4294967295
	fmt.Printf("\t Unsigned Integer 32 with range of 0 to 4294967295: %v \n", unSignedInt32)

	var unSignedInt64 uint64 = 18446744073709551615
	fmt.Printf("\t Unsigned Integer 64 with range of 0 to 18446744073709551615: %v \n", unSignedInt64)
	fmt.Println("\t ====================================================")

	var float32 float32 = 18.1234567
	fmt.Printf("\t Float 32 with max value of 7 decimal digits: %v \n", float32)

	var float64 float64 = 18.12345678910111213
	fmt.Printf("\t Float 64 with max value of 15 decimal digits: %v \n", float64)
}

func textDataType() {
	var string string = "Juan"
	fmt.Printf("\t String data type: %v \n", string)

	var character rune = 'J'
	fmt.Printf("\t Character data type: %c \n", character)
}

func dataStructureDataType() {

}

func specialDataType() {

}
