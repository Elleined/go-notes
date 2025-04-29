package main

import (
	"fmt"
	"time"
)

type MyCustomError struct {
	Message    string
	StatusCode int
	Timestamp  time.Time
}

func (e *MyCustomError) Error() string {
	return fmt.Sprintf("Error occurred: %s with status code of %d at exactly %v", e.Message, e.StatusCode, e.Timestamp)
}

func main() {
	number, err := errorSimulation()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(number)
}

func errorSimulation() (int, error) {
	if true {
		return 0, &MyCustomError{
			Message:    "Should not be null",
			StatusCode: 403,
			Timestamp:  time.Now(),
		}
	}

	return 1, nil
}
