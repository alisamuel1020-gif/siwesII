package main

import "fmt"

type Employee struct {
	ID            string
	Name          string
	Age           int
	Department    string
	MonthlySalary int
}

func increaseSalary(emp *Employee, amount int) {
	emp.MonthlySalary += amount
}

func main() {
	employee := Employee{
		ID:            "EMP001",
		Name:          "Ali Samuel",
		Age:           25,
		Department:    "Engineering",
		MonthlySalary: 250000,
	}

	fmt.Println("===== EMPLOYEE SALARY UPDATE =====")
	fmt.Println("Employee ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Department:", employee.Department)
	fmt.Println("Original Monthly Salary:", employee.MonthlySalary)

	increaseSalary(&employee, 50000)
	fmt.Println("Updated Monthly Salary:", employee.MonthlySalary)

	nilPointer()
}
