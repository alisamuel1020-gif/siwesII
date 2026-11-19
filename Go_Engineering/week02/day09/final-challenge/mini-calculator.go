package main

import "fmt"

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
func main() {
	fmt.Println("Addition:", add(20.0, 5.0))
	fmt.Println("Subtraction:", subtract(20.0, 5.0))
	fmt.Println("Multiplication:", multiply(20.0, 5.0))

	result1, err := divide(20.0, 5.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", result1)
	}

	result2, err := divide(20.0, 0.0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", result2)
	}

}
