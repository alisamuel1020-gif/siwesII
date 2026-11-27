package main

import (
	"errors"
	"fmt"
)

func validateAge(age int) error {
	if age < 18 {
		return errors.New("employee must be at least 18 years old")
	}

	return nil
}

func main() {
	err := validateAge(16)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Age validation successful")
}
