package main

import "fmt"

type Employee struct {
	ID            string
	Name          string
	Age           int
	MonthlySalary int
}

func getEmployeeDetails() (string, string, int, int) {
	employeeID := "EMP001"
	employeeName := "Ali Samuel"
	Age := 25
	MonthlySalary := 250000

	return employeeID, employeeName, Age, MonthlySalary
}

func calculateBonus(salary int) (int, error) {
	if salary <= 0 {
		return 0, fmt.Errorf("salary must be greater than zero")
	}

	bonus := salary * 10 / 100

	return bonus, nil
}

func main() {
	employeeID, employeeName, age, monthlySalary := getEmployeeDetails()
	emp := Employee{
		ID:            employeeID,
		Name:          employeeName,
		Age:           age,
		MonthlySalary: monthlySalary,
	}

	bonus, err := calculateBonus(emp.MonthlySalary)

	if err != nil {
		fmt.Println("Error calculating bonus:", err)
		return
	}

	fmt.Println("====== Employee Details ======")
	fmt.Println("ID:            ", emp.ID)
	fmt.Println("Name:          ", emp.Name)
	fmt.Println("Age:           ", emp.Age)
	fmt.Println("Monthly Salary:", emp.MonthlySalary)
	fmt.Println("Bonus (10%):   ", bonus)

}
