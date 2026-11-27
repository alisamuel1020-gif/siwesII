package main

import "fmt"

type Employee struct {
	ID            string
	Name          string
	Department    string
	MonthlySalary int
}

func calculateSalary(basicSalary int, allowance int) int {
	total := basicSalary + allowance
	return total
}

func getSalaryDetails(monthlySalary int) (int, int) {
	annualSalary := monthlySalary * 12
	return monthlySalary, annualSalary
}

func calculateAverageSalary(salaries ...int) float64 {
	if len(salaries) == 0 {
		return 0
	}

	total := 0

	for _, salary := range salaries {
		total += salary
	}

	average := float64(total) / float64(len(salaries))
	return average
}

func createSalaryIncrement(currentSalary int) func() int {
	incrementedSalary := currentSalary

	return func() int {
		incrementedSalary += 20000
		return incrementedSalary
	}
}

func main() {
	employee := Employee{
		ID:            "EMP001",
		Name:          "Ali Samuel",
		Department:    "Engineering",
		MonthlySalary: 250000,
	}

	totalSalary := calculateSalary(250000, 50000)

	employee.MonthlySalary = totalSalary

	monthly, annual := getSalaryDetails(employee.MonthlySalary)

	fmt.Printf("Monthly Salary: %d\n Annual Salary: %d\n", monthly, annual)

	fmt.Println("Basic Salary:", 250000)
	fmt.Println("Allowance:", 50000)
	fmt.Println("Total Monthly Salary:", totalSalary)

	fmt.Println("Employee ID:", employee.ID)
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Department:", employee.Department)
	fmt.Println("Monthly Salary:", employee.MonthlySalary)

	average := calculateAverageSalary(300000, 250000, 350000)
	fmt.Println("Average Employee Salary:", average)

	calculateBonus := func(monthlySalary int) int {
		bonus := float64(monthlySalary) * 0.1
		return int(bonus)
	}

	bonus := calculateBonus(employee.MonthlySalary)

	fmt.Println("Employee Bonus:", bonus)

	increment := createSalaryIncrement(employee.MonthlySalary)

	for i := 0; i < 3; i++ {
		employee.MonthlySalary = increment()
		fmt.Println("Updated Salary:", employee.MonthlySalary)
	}
}
