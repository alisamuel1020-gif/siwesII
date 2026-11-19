package main

import "fmt"

func main() {
	studentName := "Winston Paul"

	// SCORES OF FIVE SUBJECTS
	math221 := 40
	math232 := 50
	math212 := 60
	math222 := 70
	sta203 := 80

	//CALCULATING TOTAL SCORE
	totalScore := math221 + math232 + math212 + math222 + sta203

	//CALCULATING SCORES
	averageScore := float64(totalScore) / 5
	targetAverage := 75
	requiredTotal := targetAverage * 5
	additionalMarksNeeded := requiredTotal - totalScore
	// DETERMINING IF STUDENT PASS OR NOT

	passed := averageScore >= 40
	distinction := averageScore >= 75
	totalIsEven := totalScore%2 == 0
	// PRINTING RESULTS
	fmt.Printf("Student Name: %s\n", studentName)
	fmt.Printf("Total Score: %d\n", totalScore)
	fmt.Printf("Average Score: %.2f\n", averageScore)
	fmt.Printf("Passed: %t\n", passed)
	fmt.Printf("Distinction: %t\n", distinction)
	fmt.Printf("Total Score is Even: %t\n", totalIsEven)
	fmt.Printf("Target Average: %d\n", targetAverage)
	fmt.Printf("Required Total for Target Average: %d\n", requiredTotal)
	fmt.Printf("Additional Marks Needed: %d\n", additionalMarksNeeded)

}
