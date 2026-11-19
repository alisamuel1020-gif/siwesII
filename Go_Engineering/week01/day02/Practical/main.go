package main

import "fmt"

func main() {
	fmt.Println("===== RESULT PROCESSING SYSTEM =====")
	// STUDENTS DETAILS
	fullName := "Jonny Cage"
	regNumber := "22/76543/edu"
	age := 20
	department := "Mathematics Education"

	// SUBJECTS AND SCOREES
	const passMark int = 40
	const maxScore int = 100

	math221 := 40
	math232 := 50
	math212 := 60
	math222 := 70
	edu203 := 80

	totalScore := math212 + math221 + math222 + math232 + edu203

	averageScore := float64(totalScore) / 5

	//DISPLAYING DATAS

	fmt.Printf("Name: %s \n", fullName)
	fmt.Printf("Age : %d \n", age)
	fmt.Printf("Registration Number: %s \n", regNumber)
	fmt.Printf("Department: %s \n", department)
	fmt.Println("_______________________")
	fmt.Println("== Subject == Score ==")
	fmt.Println("_______________________")
	fmt.Printf("   math221       %d \n", math221)
	fmt.Printf("   math232       %d \n", math232)
	fmt.Printf("   math212       %d \n", math212)
	fmt.Printf("   math222       %d \n", math222)
	fmt.Printf("   edu 203       %d \n", edu203)
	fmt.Println("===================================")
	fmt.Printf("maximum examination score = %d \n", maxScore)
	fmt.Printf("Total Score = %d \n", totalScore)
	fmt.Printf("Averge Score = %.2f \n", averageScore)
	fmt.Println("===================================")

}
