package main

import "fmt"

type Employee struct {
	ID            string
	Name          string
	Age           int
	Department    string
	MonthlySalary int
}

func main() {

	employees := map[string]string{
		"EMP001": "Ali Samuel",
		"EMP002": "John Peter",
		"EMP003": "Mary James",
	}

	// Add a new employee
	employees["EMP004"] = "David James"

	// Access an employee
	fmt.Println("Employee EMP001:", employees["EMP001"])

	// Update an employee
	employees["EMP002"] = "John Samuel"

	// Display the updated map
	fmt.Println("Updated Employees:", employees)

	employeeName, exists := employees["EMP005"]

	if exists {
		fmt.Println("Employee found:", employeeName)
	} else {
		fmt.Println("Employee not found")
	}

	// Delete an employee
	delete(employees, "EMP003")

	// Display employees after deletion
	fmt.Println("Employees after deletion:", employees)

	fmt.Println("\n===== EMPLOYEE RECORDS =====")

	for employeeID, employeeName := range employees {
		fmt.Printf("Employee ID: %s | Employee Name: %s\n", employeeID, employeeName)
	}

	employeeAges := map[string]int{
		"EMP001": 25,
		"EMP002": 30,
		"EMP004": 28,
	}

	// Add a new employee age
	employeeAges["EMP005"] = 35

	// Display employee ages
	fmt.Println("\n===== EMPLOYEE AGES =====")

	for employeeID, age := range employeeAges {
		fmt.Printf("Employee ID: %s | Age: %d\n", employeeID, age)
	}

	employeeRecords := map[string]Employee{
		"EMP001": {
			ID:            "EMP001",
			Name:          "Ali Samuel",
			Age:           25,
			Department:    "Engineering",
			MonthlySalary: 250000,
		},
		"EMP002": {
			ID:            "EMP002",
			Name:          "John Peter",
			Age:           30,
			Department:    "Human Resources",
			MonthlySalary: 300000,
		},
	}

	// Retrieve EMP001
	employee := employeeRecords["EMP001"]

	// Update employee information
	employee.Age = 26
	employee.MonthlySalary = 300000

	// Store the modified employee back in the map
	employeeRecords["EMP001"] = employee

	// Display updated employee
	fmt.Println("\n===== UPDATED EMPLOYEE =====")
	fmt.Printf("Name: %s\n", employeeRecords["EMP001"].Name)
	fmt.Printf("Age: %d\n", employeeRecords["EMP001"].Age)
	fmt.Printf("Monthly Salary: %d\n", employeeRecords["EMP001"].MonthlySalary)

	fmt.Println("\n===== EMPLOYEE RECORDS =====")

	for employeeID, employee := range employeeRecords {
		fmt.Printf("\nEmployee ID: %s\n", employeeID)
		fmt.Printf("Name: %s\n", employee.Name)
		fmt.Printf("Age: %d\n", employee.Age)
		fmt.Printf("Department: %s\n", employee.Department)
		fmt.Printf("Monthly Salary: %d\n", employee.MonthlySalary)
	}

}
