package main

import "fmt"

type Employee struct {
	ID            string
	Name          string
	Age           int
	Department    string
	MonthlySalary int
}

var employees = make(map[string]Employee)

func main() {
	fmt.Println("===== EMPLOYEE RECORDS MANAGEMENT SYSTEM =====")

	registerEmployee(Employee{
		ID:            "EMP001",
		Name:          "Ali Samuel",
		Age:           25,
		Department:    "Engineering",
		MonthlySalary: 250000,
	})

	registerEmployee(Employee{
		ID:            "EMP002",
		Name:          "John Peter",
		Age:           30,
		Department:    "Human Resources",
		MonthlySalary: 300000,
	})

	registerEmployee(Employee{
		ID:            "EMP001",
		Name:          "Samuel James",
		Age:           28,
		Department:    "Finance",
		MonthlySalary: 350000,
	})

	fmt.Println("\nTotal Employees:", len(employees))

	searchEmployee("EMP001")

	updateEmployee("EMP001", "Cloud Engineering", 400000)

	deleteEmployee("EMP001")

	displayAllEmployees()

	fmt.Println("\nTotal Employees:", len(employees))
}

func registerEmployee(emp Employee) {
	if _, exists := employees[emp.ID]; exists {
		fmt.Println("Employee ID already exists!")
		return
	}

	employees[emp.ID] = emp

	fmt.Println("Employee registered successfully!")
}

func searchEmployee(employeeID string) {
	employee, exists := employees[employeeID]

	if !exists {
		fmt.Println("Employee not found!")
		return
	}

	fmt.Println("\n===== EMPLOYEE DETAILS =====")
	fmt.Println("Employee ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Department:", employee.Department)
	fmt.Println("Monthly Salary:", employee.MonthlySalary)
}

func updateEmployee(employeeID string, newDepartment string, newSalary int) {
	employee, exists := employees[employeeID]

	if !exists {
		fmt.Println("Employee not found!")
		return
	}

	employee.Department = newDepartment
	employee.MonthlySalary = newSalary

	employees[employeeID] = employee

	fmt.Println("Employee record updated successfully!")
}

func deleteEmployee(employeeID string) {
	if _, exists := employees[employeeID]; !exists {
		fmt.Println("Employee not found!")
		return
	}

	delete(employees, employeeID)

	fmt.Println("Employee deleted successfully!")
}

func displayAllEmployees() {
	if len(employees) == 0 {
		fmt.Println("No employee records found!")
		return
	}

	fmt.Println("\n===== ALL EMPLOYEE RECORDS =====")

	for employeeID, employee := range employees {
		fmt.Println("------------------------------")
		fmt.Println("Employee ID:", employeeID)
		fmt.Println("Name:", employee.Name)
		fmt.Println("Age:", employee.Age)
		fmt.Println("Department:", employee.Department)
		fmt.Println("Monthly Salary:", employee.MonthlySalary)
	}
}
