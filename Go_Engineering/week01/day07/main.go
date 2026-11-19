package main

import "fmt"

func main() {
	//Student Information Management System

	fmt.Println("====== Student Information Management System ======")

	// Creating an empty map for student information
	studentInfo := make(map[string]string)

	// Storing a student's name
	studentInfo["name"] = "Ede Winifred"

	// Storing their department
	studentInfo["department"] = "Computer Science"

	// Storing their university
	studentInfo["university"] = "SimplyTech University"

	// Storing their level
	studentInfo["level"] = "Final Year"

	// Displaying all student information
	fmt.Println("Student Information:")
	for key, value := range studentInfo {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Update the student's level
	studentInfo["level"] = "Graduate"
	fmt.Println("\nUpdated Level:", studentInfo["level"])

	// Check whether a particular key exists
	keyToCheck := "department"
	if value, exists := studentInfo[keyToCheck]; exists {
		fmt.Printf("\nKey '%s' exists with value: %s\n", keyToCheck, value)
	} else {
		fmt.Printf("\nKey '%s' does not exist.\n", keyToCheck)
	}

	// Delete a particular entry
	keyToDelete := "university"
	delete(studentInfo, keyToDelete)
	fmt.Printf("\nDeleted key '%s'.\n", keyToDelete)

	// Display the total number of remaining records
	fmt.Printf("\nTotal number of remaining records: %d\n", len(studentInfo))

	//creating a nested map for student information
	studentRecords := make(map[string]map[string]string)

	// Adding a student record
	studentRecords["student1"] = map[string]string{
		"name":       "Ede Winifred",
		"department": "Computer Science",
		"university": "SimplyTech University",
		"level":      "Final Year",
	}

	// NESTED MAP: STUDENT SUBJECT SCORES
	students := map[string]map[string]int{
		"Ali": {
			"Mathematics": 80,
			"English":     90,
			"Physics":     75,
		},
		"John": {
			"Mathematics": 85,
			"English":     70,
			"Physics":     80,
		},
		"Peter": {
			"Mathematics": 95,
			"English":     88,
			"Physics":     90,
		},
	}
	fmt.Println("\nAll Students and Their Scores:")

	for studentName, subjects := range students {
		fmt.Println("\nStudent:", studentName)

		for subject, score := range subjects {
			fmt.Printf("%s: %d\n", subject, score)
		}
	}
	fmt.Println("\nAli's Mathematics Score:", students["Ali"]["Mathematics"])
	students["John"]["English"] = 85

	total := 0

	for _, score := range students["Ali"] {
		total += score
	}

	fmt.Println("\nTotal Score for Ali:", total)

	fmt.Println("\nNested Student Records:")
	for studentID, info := range studentRecords {
		fmt.Printf("Student ID: %s\n", studentID)
		for key, value := range info {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
	fmt.Println("\nRemaining Student Information:")

	for key, value := range studentInfo {
		fmt.Printf("%s: %s\n", key, value)
	}
}
