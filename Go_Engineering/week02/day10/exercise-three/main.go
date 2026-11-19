package main

import "fmt"

var companyName = "simplyTech International"

func main() {
	employeeName := "John Bush"
	salary := 250000

	displayEmployee := func() {
		fmt.Println("DISPLAYING EMPLOYEE DETA")
		fmt.Printf("Name: %s\n Salary: %d \n  Company: %s\n", employeeName, salary, companyName)

	}

	increaseSalary := func() {
		salary += 15000
		fmt.Println("New Salary = ", salary)
	}
	displayEmployee()

	increaseSalary()
	displayEmployee()

	increaseSalary()
	displayEmployee()

}
