package main

import (
	"errors"
	"fmt"
	"strings"
)

type Employee struct {
	ID            string
	Name          string
	Age           int
	MonthlySalary int
}

func validateAge(age int) error {
	if age < 18 {
		return errors.New("employee must be at least 18 years old")
	}

	if age > 65 {
		return fmt.Errorf("invalid age: %d. Maximum age is 65", age)
	}

	return nil
}

func validateSalary(salary int) error {
	if salary <= 0 {
		return fmt.Errorf("invalid monthly salary: %d. Salary must be greater than zero", salary)
	}

	return nil
}

func validateEmployeeID(id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("employee ID cannot be empty")
	}
	return nil
}

func main() {
	emp := Employee{
		ID:            "EMP001",
		Name:          "Ali Samuel",
		Age:           25,
		MonthlySalary: 250000,
	}

	err := validateAge(emp.Age)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = validateSalary(emp.MonthlySalary)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = validateEmployeeID(emp.ID)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("====== EMPLOYEE VALIDATION SYSTEM ======")
	fmt.Println("Employee ID:", emp.ID)
	fmt.Println("Name:", emp.Name)
	fmt.Println("Age:", emp.Age)
	fmt.Println("Monthly Salary:", emp.MonthlySalary)
}
