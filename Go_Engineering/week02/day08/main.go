package main

import "fmt"

// EXERCISE ONE -- GREETING FUNCTION

func greetUser(name string) {
	fmt.Printf("Hello!, %s!\n", name)
}

// EXERCISE TWO -- STUDENT INFORMATION

func displayStudent(studentName string, department string, level string) {
	fmt.Printf("Student: %s\nDepartment: %s\nLevel: %s\n", studentName, department, level)
}

// EXERCISE THREE -- CALCULATOR

func add(a, b int) {
	fmt.Println("The sum of ", a, "and", b, "is", a+b)
}

func subtract(a, b int) {
	fmt.Println("The difference between ", a, "and", b, "is", a-b)
}

func multiply(a, b int) {
	fmt.Println("The product of ", a, "and", b, "is", a*b)
}

func divide(a, b int) {
	if b != 0 {
		fmt.Println("The quotient of ", a, "and", b, "is", a/b)
	} else {
		fmt.Println("Error: Division by zero is not allowed.")
	}
}
func main() {
	// EXERCISE ONE -- GREETING FUNCTION
	greetUser("John")
	greetUser("James")
	greetUser("Paul")

	// EXERCISE TWO -- STUDENT INFORMATION
	displayStudent("John", "Computer Science", "400L")
	displayStudent("Oyale", "HSM", "400L")
	displayStudent("Oche", "Computer Science", "400L")

	// EXERCISE THREE -- CALCULATOR
	add(20, 5)
	add(100, 25)
	add(15, 3)
	subtract(20, 5)
	subtract(100, 25)
	subtract(15, 3)
	multiply(20, 5)
	multiply(100, 25)
	multiply(15, 3)
	divide(20, 5)
	divide(100, 25)
	divide(15, 3)
}
