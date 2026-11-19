package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	ID            string
	Name          string
	Age           int
	Department    string
	MonthlySalary int
}

func employeeRegistration(reader *bufio.Reader, existingEmployees []Employee) (Employee, bool) {
	var emp Employee

	fmt.Println("===== EMPLOYEE REGISTRATION SYSTEM =====")

	fmt.Print("Enter Employee ID: ")
	employeeID, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading employee ID:", err)
		return Employee{}, false
	}

	employeeID = strings.TrimSpace(employeeID)
	if employeeID == "" {
		fmt.Println("Employee ID cannot be empty")
		return Employee{}, false
	}

	if !strings.HasPrefix(employeeID, "EMP") {
		fmt.Println("Employee ID must start with \"EMP\"")
		return Employee{}, false
	}

	for _, existing := range existingEmployees {
		if existing.ID == employeeID {
			fmt.Printf("Employee ID %q is already registered\n", employeeID)
			return Employee{}, false
		}
	}

	emp.ID = employeeID

	fmt.Print("Enter Employee Name: ")
	employeeName, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading employee Name:", err)
		return Employee{}, false
	}
	employeeName = strings.TrimSpace(employeeName)
	if employeeName == "" {
		fmt.Println("Employee name cannot be empty")
		return Employee{}, false
	}

	emp.Name = employeeName

	fmt.Print("Enter Employee Department: ")

	employeeDepartment, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading employee Department:", err)
		return Employee{}, false
	}

	employeeDepartment = strings.TrimSpace(employeeDepartment)

	if employeeDepartment == "" {
		fmt.Println("Employee department cannot be empty")
		return Employee{}, false
	}

	emp.Department = employeeDepartment

	fmt.Print("Enter Your Age: ")
	stringEmployeeAge, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading employee Age:", err)
		return Employee{}, false
	}

	stringEmployeeAge = strings.TrimSpace(stringEmployeeAge)
	employeeAge, err := strconv.Atoi(stringEmployeeAge)
	if err != nil {
		fmt.Println("Invalid User Age")
		return Employee{}, false
	}

	if employeeAge < 18 || employeeAge > 65 {
		fmt.Println("Please enter a valid age")
		return Employee{}, false
	}
	emp.Age = employeeAge

	fmt.Print("Enter Your Salary: ")

	stringEmployeeMonthlySalary, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading employee Salary:", err)
		return Employee{}, false
	}

	stringEmployeeMonthlySalary = strings.TrimSpace(stringEmployeeMonthlySalary)
	employeeMonthlySalary, err := strconv.Atoi(stringEmployeeMonthlySalary)
	if err != nil {
		fmt.Println("Invalid User Monthly Salary")
		return Employee{}, false
	}
	if employeeMonthlySalary <= 0 {
		fmt.Println("Monthly salary must be greater than zero")
		return Employee{}, false
	}

	emp.MonthlySalary = employeeMonthlySalary

	return emp, true

}

func displayEmployee(emp Employee) {

	annualSalary := emp.MonthlySalary * 12

	fmt.Println("===== EMPLOYEE DETAILS =====")

	fmt.Printf("Employee ID: %s\nEmployee Name: %s\nEmployee Age: %d\nEmployee Department: %s\nEmployee Monthly Salary: %d\nAnnual Salary: %d\n", emp.ID, emp.Name, emp.Age, emp.Department, emp.MonthlySalary, annualSalary)
}

func updateEmployeeSalary(emp *Employee) {
	emp.MonthlySalary = 750000
}

func changeAgeByValue(emp Employee) {
	emp.Age = 40
}

func changeAgeByPointer(emp *Employee) {
	emp.Age = 50
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	emp001 := Employee{
		ID:            "EMP001",
		Name:          "John",
		Age:           25,
		Department:    "Engineering",
		MonthlySalary: 2000000,
	}

	employees := []Employee{
		{
			ID:            "EMP001",
			Name:          "John",
			Age:           25,
			Department:    "Engineering",
			MonthlySalary: 2000000,
		},
		{
			ID:            "EMP002",
			Name:          "Samuel",
			Age:           24,
			Department:    "Software Development",
			MonthlySalary: 3000000,
		},
	}

	employees = append(employees, Employee{
		ID:            "EMP003",
		Name:          "Davide",
		Age:           28,
		Department:    "Cybersecurity",
		MonthlySalary: 2500000,
	})

	for {
		newEmployee, success := employeeRegistration(reader, employees)
		fmt.Print("Register another employee? (yes/no): ")
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading response:", err)
			break
		}

		if success {
			employees = append(employees, newEmployee)
		}

		userInput = strings.TrimSpace(userInput)
		if userInput == "no" {
			fmt.Println("Thank you for using the Employee Registration System!")
			break
		}

	}

	for _, emp := range employees {
		displayEmployee(emp)
	}

	changeAgeByValue(emp001)
	fmt.Println(emp001.Age)

	changeAgeByPointer(&emp001)
	fmt.Println(emp001.Age)

	updateEmployeeSalary(&emp001)

	displayEmployee(emp001)
}
