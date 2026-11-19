package main

import "fmt"

type Employee struct {
	ID            string
	Name          string
	Age           int
	Department    string
	MonthlySalary int
	AnnualSalary  int
}

func (e Employee) displayEmployeeDetails() {
	fmt.Println("Employee ID:", e.ID)
	fmt.Println("Employee Name:", e.Name)
	fmt.Println("Employee Age:", e.Age)
	fmt.Println("Employee Department:", e.Department)
	fmt.Println("Employee Monthly Salary:", e.MonthlySalary)
	fmt.Println("Employee Annual Salary:", e.AnnualSalary)
}

func (e Employee) calculateAnnualSalary() int {
	return e.MonthlySalary * 12
}

func (e *Employee) increaseSalary() {
	e.MonthlySalary += 50000
	fmt.Println("Salary inside method:", e.MonthlySalary)
}

func (e *Employee) updateAge(newAge int) {
	e.Age = newAge
}

func main() {
	employee1 := Employee{
		ID:            "EMP001",
		Name:          "Samuel",
		Age:           23,
		Department:    "Engineering",
		MonthlySalary: 350000,
	}

	employee2 := Employee{
		ID:            "EMP002",
		Name:          "Samson",
		Age:           27,
		Department:    "Telecom",
		MonthlySalary: 300000,
	}

	// Calculate initial annual salaries
	employee1.AnnualSalary = employee1.calculateAnnualSalary()
	employee2.AnnualSalary = employee2.calculateAnnualSalary()

	// Increase Samuel's salary
	employee1.increaseSalary()

	// Recalculate Samuel's annual salary
	employee1.AnnualSalary = employee1.calculateAnnualSalary()

	// Increasing age
	employee1.updateAge(40)
	employee2.updateAge(32)

	// Display employee details
	employee1.displayEmployeeDetails()
	employee2.displayEmployeeDetails()

}
