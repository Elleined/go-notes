package main

import "fmt"

func main() {
	variableTopic()
	dataTypesTopic()
	decisionMakingAndControlFlowTopic()
}

func variableTopic() {
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

func dataTypesTopic() {
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
}

func decisionMakingAndControlFlowTopic() {
	fmt.Println("\n\n=========================")
	fmt.Println("DECISION MAKING")
	fmt.Println("=========================")
	decisionMaking()
	fmt.Println("================== END DECISION MAKING ==================")

	fmt.Println("\n\n=========================")
	fmt.Println("CONTROL FLOW")
	fmt.Println("=========================")
	controlFlow()
	fmt.Println("================== END CONTROL FLOW ==================")
}

func longHandAndShortHandVariableDeclaration() {
	fmt.Print("\t === long hand === \n")
	var longHand string = "Denielle"
	fmt.Printf("\t long hand declaration: %s \n", longHand)

	fmt.Print("\t === short hand === \n")
	shortHand := "Denielle"
	fmt.Printf("\t short hand declaration: %s \n", shortHand)
}

func varAndConst() {
	fmt.Print("\t === var === \n")
	var name string = "Red"
	fmt.Printf("\t Variable created via var keyword: %s \n", name)

	fmt.Print("\t === const === \n")
	const AGE uint8 = 22
	fmt.Printf("\t Variable created via const keyword: %d \n", AGE)
}

func numericDataType() {
	fmt.Println("\t === Integer ===")
	var signedInt8 int8 = 127
	fmt.Printf("\t Integer 8 with range of -128 to 127: %v \n", signedInt8)

	var signedInt16 int16 = 32767
	fmt.Printf("\t Integer 16 with range of -32,768 to 32767: %v \n", signedInt16)

	var signedInt32 int32 = 2147483647
	fmt.Printf("\t Integer 32 with range of -2,147,483,648 to 2147483647: %v \n", signedInt32)

	var signedInt64 int64 = 9223372036854775807
	fmt.Printf("\t Integer 64 with range of -9,223,372,036,854,775,808 to 9223372036854775807: %v \n", signedInt64)

	fmt.Println("\t === Unsigned Integer ===")
	var unSignedInt8 uint8 = 255
	fmt.Printf("\t Unsigned Integer 8 with range of 0 to 255: %v \n", unSignedInt8)

	var unSignedInt16 uint16 = 65535
	fmt.Printf("\t Unsigned Integer 16 with range of 0 to 65535: %v \n", unSignedInt16)

	var unSignedInt32 uint32 = 4294967295
	fmt.Printf("\t Unsigned Integer 32 with range of 0 to 4294967295: %v \n", unSignedInt32)

	var unSignedInt64 uint64 = 18446744073709551615
	fmt.Printf("\t Unsigned Integer 64 with range of 0 to 18446744073709551615: %v \n", unSignedInt64)

	fmt.Println("\t === Float ===")
	var float32 float32 = 18.1234567
	fmt.Printf("\t Float 32 with max value of 7 decimal digits: %v \n", float32)

	var float64 float64 = 18.12345678910111213
	fmt.Printf("\t Float 64 with max value of 15 decimal digits: %v \n", float64)
}

func textDataType() {
	fmt.Println("\t === String ===")
	var string string = "Juan"
	fmt.Printf("\t String data type: %v \n", string)

	fmt.Println("\t === Character ===")
	var character rune = 'J'
	fmt.Printf("\t Character data type: %c \n", character)
}

func dataStructureDataType() {
	names := []string{
		"Red",
		"Blue",
	}
	fmt.Println("\t === Slice ===")
	for i, v := range names {
		fmt.Printf("\t Index: %v, Value: %v \n", i, v)
	}

	persons := map[int]string{
		1: "Red",
		2: "Blue",
	}
	fmt.Println("\t === Map ===")
	for k, v := range persons {
		fmt.Printf("\t Key: %d Value: %s \n", k, v)
	}

	fmt.Println("\t === Struct ===")
	type Person struct {
		id   uint
		Name string
	}
	person := Person{
		1,
		"Red",
	}
	fmt.Printf("\t %+v \n", person)
}

func decisionMaking() {
	fmt.Println("\t === If ===")
	const age uint8 = 15
	if age < 20 {
		fmt.Println("\t Executed because age is less than 20")
	} else {
		fmt.Println("\t Executed because age is greater than 20")
	}

	fmt.Println("\t === If Else ===")
	if age < 20 {
		fmt.Println("\t Executed because age is less than 20")
	} else if age < 25 {
		fmt.Println("\t Executed because age is less than 25")
	} else {
		fmt.Println("\t Executed because age is more than 25")
	}

	fmt.Println("\t === IF WITH SHORT STATEMENT ===")
	// Here score is only available within the if statement scope and cannot be access outside the if statement.
	if score := 10; score > 10 {
		fmt.Printf("\t Score is greater than %d", score)
	}
}

func controlFlow() {
	fmt.Println("\t === FOR I ===")
	for i := 1; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	fmt.Println("\n \t === FOR EACH ===")
	numbers := []int{
		1,
		2,
		3,
	}
	for _, number := range numbers {
		fmt.Print(number)
	}
	fmt.Println()
}
