package main

import (
	"errors"
	"fmt"
)

// This error.go contains the ff:
// 1. errors.New
// 2. errors.Is
// 3. errors.As
// 4. errors.Join

func main() {
	num, err := errorNewFunc()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(num)
}

func errorNewFunc() (int, error) {
	if true {
		return 0, errors.New("Error occured!")
	}
	return 1, nil
}
