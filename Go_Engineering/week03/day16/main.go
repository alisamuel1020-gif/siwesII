package main

import "fmt"

func calculateSalary(basicSalary int, allowance int) int {
	total := basicSalary + allowance
	return total
}

func getEmployeeDetails() (string, string, int) {
	ID := "EMP001"
	name := "Ali Samuel"
	salary := 250000

	return ID, name, salary
}

func calculateAverageSalary(salaries ...int) float64 {
	total := 0

	if len(salaries) == 0 {
		return 0
	}

	for _, salary := range salaries {
		total += salary
	}

	average := float64(total) / float64(len(salaries))

	return average
}

func salaryIncrement() func() int {
	salary := 300000

	return func() int {
		salary += 20000
		return salary
	}
}

func main() {
	salary := calculateSalary(250000, 50000)

	fmt.Println("Total Salary", salary)

	employeeID, employeeName, monthlySalary := getEmployeeDetails()

	fmt.Println("Employee ID:", employeeID)
	fmt.Println("Employee Name:", employeeName)
	fmt.Println("Monthly Salary:", monthlySalary)

	average := calculateAverageSalary(250000, 300000, 350000)

	fmt.Println("Average Salary:", average)

	calculateBonus := func(salary int) int {
		bonus := float64(salary) * 0.1

		return int(bonus)
	}

	bonus := calculateBonus(250000)

	fmt.Println("Employee Bonus:", bonus)

	increment := salaryIncrement()

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
