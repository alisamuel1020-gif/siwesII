package main

import "fmt"

func main() {
	studentName := "Greg Patrick"

	// SCORES OF FIVE SUBJECTS
	cmp221Score := 40
	cmp232Score := 50
	cmp212Score := 60
	cmp222Score := 70
	sta203Score := 80

	// SCORE VALIDATION
	if cmp221Score < 0 || cmp221Score > 100 {
		fmt.Println("Invalid score for CMP221")
		return
	}

	if cmp232Score < 0 || cmp232Score > 100 {
		fmt.Println("Invalid score for CMP232")
		return
	}

	if cmp212Score < 0 || cmp212Score > 100 {
		fmt.Println("Invalid score for CMP212")
		return
	}

	if cmp222Score < 0 || cmp222Score > 100 {
		fmt.Println("Invalid score for CMP222")
		return
	}

	if sta203Score < 0 || sta203Score > 100 {
		fmt.Println("Invalid score for STA203")
		return
	}

	// CALCULATING TOTAL AND AVERAGE
	totalScore := cmp221Score + cmp232Score + cmp212Score + cmp222Score + sta203Score
	averageScore := float64(totalScore) / 5

	// GRADE ASSIGNMENT
	var grade string
	var remark string

	if averageScore >= 75 {
		grade = "A"
		remark = "Passed with Distinction"
	} else if averageScore >= 65 {
		grade = "B"
		remark = "Good"
	} else if averageScore >= 55 {
		grade = "C"
		remark = "Satisfactory"
	} else if averageScore >= 45 {
		grade = "D"
		remark = "Fair"
	} else {
		grade = "F"
		remark = "Failed"
	}

	// PERFORMANCE FEEDBACK
	switch grade {
	case "A":
		fmt.Println("Outstanding performance")
	case "B":
		fmt.Println("Very good performance")
	case "C":
		fmt.Println("Good performance")
	case "D":
		fmt.Println("Needs improvement")
	case "F":
		fmt.Println("Failed")
	default:
		fmt.Println("Unknown grade")
	}

	// DISPLAY RESULTS
	fmt.Printf("Student Name: %s\n", studentName)
	fmt.Printf("Total Score: %d\n", totalScore)
	fmt.Printf("Average Score: %.2f\n", averageScore)
	fmt.Printf("Grade: %s\n", grade)
	fmt.Printf("Remark: %s\n", remark)
}
