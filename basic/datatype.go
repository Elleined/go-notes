package main

import "fmt"

func main() {
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
