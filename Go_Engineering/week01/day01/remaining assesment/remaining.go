package main

import "fmt"

func main() {
	fmt.Println("================================")
	fmt.Println("      EMPLOYEE INFORMATION")
	fmt.Println("================================")

	// Employee One
	employeeOneName := "George D"
	employeeOneDepartment := "Engineering"
	employeeOneJobRole := "Software Engineer"
	employeeOneID := "001"

	// Employee Two
	employeeTwoName := "Benson Paul"
	employeeTwoDepartment := "Engineering"
	employeeTwoJobRole := "Software Engineer"
	employeeTwoID := "002"

	// Employee Three
	employeeThreeName := "Sammy Curry"
	employeeThreeDepartment := "Engineering"
	employeeThreeJobRole := "Software Engineer"
	employeeThreeID := "003"

	fmt.Println("\nEmployee One")
	fmt.Printf("ID: %s\nName: %s\nDepartment: %s\nJob Role: %s\n",
		employeeOneID, employeeOneName,
		employeeOneDepartment, employeeOneJobRole)

	fmt.Println("\nEmployee Two")
	fmt.Printf("ID: %s\nName: %s\nDepartment: %s\nJob Role: %s\n",
		employeeTwoID, employeeTwoName,
		employeeTwoDepartment, employeeTwoJobRole)

	fmt.Println("\nEmployee Three")
	fmt.Printf("ID: %s\nName: %s\nDepartment: %s\nJob Role: %s\n",
		employeeThreeID, employeeThreeName,
		employeeThreeDepartment, employeeThreeJobRole)

	fmt.Println("================================")
}
