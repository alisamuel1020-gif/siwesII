package main

import "fmt"

// EXERCISE ONE -- ADDITION FUNCTION

func addNumbers(a, b int) int {
	return a + b
}

//EXERCISE TWO -- MULTIPLICATION FUNCTION

func multiplyNumbers(a, b int) int {
	return a * b
}

//FUNCTION RETURNING A STRING

func getCountry() string {
	return "Nigeria"
}

// FUNCTION RETURNING BOOLEAN

func isEligible(age int) bool {
	return age >= 18
}

// RETURNING MULTIPLE VALUES

func calculateNumbers(a, b int) (int, int) {
	//calculating
	sum := a + b
	difference := a - b

	//returning values
	return sum, difference
}

// HANDLING ERRORS

func divideNumbers(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result1 := addNumbers(20, 30)
	fmt.Println("The sum is:", result1)

	result2 := multiplyNumbers(6, 7)
	fmt.Println("The product is:", result2)

	country := getCountry()
	fmt.Println("My country is:", country)

	eligible1 := isEligible(20)
	fmt.Println("Is eligible:", eligible1)

	eligible2 := isEligible(15)
	fmt.Println("Is eligible:", eligible2)

	sum, difference := calculateNumbers(20, 10)
	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)

	outcome1, err := divideNumbers(20, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", outcome1)
	}

	outcome2, err := divideNumbers(20, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", outcome2)
	}
}
