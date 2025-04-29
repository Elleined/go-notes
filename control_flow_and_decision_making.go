package main

import "fmt"

func main() {
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

	fmt.Println("\n \t === FOR RANGE ===")
	numbers := []int{
		1,
		2,
		3,
	}
	for index, number := range numbers {
		fmt.Printf("Index: %d Number: %d", index, number)
	}
	fmt.Println()

	fmt.Println("\n \t === WHILE LOOP ===")
	counter := 10
	for counter > 0 {
		fmt.Printf("Counter: %d", counter)
		counter--
	}
	fmt.Println()

	fmt.Println("\n \t === INFINITE LOOP ===")
	for {
		fmt.Println("I'm stuck here! HELP")
	}
}
