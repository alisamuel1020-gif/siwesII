package main

import "fmt"

func main() {

    // Student information
    studentName := "Ali Samuel"
    registrationNumber := "JSTU/CSC/2026/001"
    age := 24
    department := "Computer Science"

    // Examination scores
    const maximumScore = 100

    mathematics := 85.0
    programming := 90.0
    database := 78.0

    // Calculate total and average
    totalScore := mathematics + programming + database
    averageScore := totalScore / 3

    // Display student information
    fmt.Println("====== STUDENT INFORMATION ======")

    fmt.Printf("Name: %s\n", studentName)
    fmt.Printf("Registration Number: %s\n", registrationNumber)
    fmt.Printf("Age: %d\n", age)
    fmt.Printf("Department: %s\n", department)

    // Display results
    fmt.Println("\n====== STUDENT RESULTS ======")

    fmt.Printf("Mathematics: %.2f\n", mathematics)
    fmt.Printf("Programming: %.2f\n", programming)
    fmt.Printf("Database: %.2f\n", database)

    fmt.Printf("Total Score: %.2f\n", totalScore)
    fmt.Printf("Average Score: %.2f\n", averageScore)

    fmt.Printf("Maximum Score Per Subject: %d\n", maximumScore)
}