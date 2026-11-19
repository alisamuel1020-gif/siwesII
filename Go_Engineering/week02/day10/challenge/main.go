package main

import "fmt"

//Mini Employee Information System

// Part A
var companyName = "Tech Solutions"

// Part B
func displayEmployee() {
	name := "Patience"
	age := 22
	salary := 250000
	fmt.Printf("Name: %s \n Age: %d \n Salary: %d \n Company Name: %s \n", name, age, salary, companyName)

	// Part C
	if age >= 18 {
		status := "Eligible"
		fmt.Println(status)
	}
}

func main() {

	func() {
		fmt.Println("Employee Management System Started!")
	}()

	displayEmployee()

}
