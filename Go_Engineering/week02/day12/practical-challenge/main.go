package main

import "fmt"

func calculateTotalScore(scores []int) int {
	totalScore := 0

	for _, score := range scores {
		totalScore += score
	}

	return totalScore
}

func calculateAverageScore(totalScore int, numberOfTests int) float64 {
	if numberOfTests == 0 {
		return 0
	}
	return float64(totalScore) / float64(numberOfTests)
}

func employeePerformanceTracker() {
	employeeTestScores := [][]int{
		{80, 85, 90},
		{70, 75, 80},
		{90, 95, 100},
	}

	for employeeIndex, scores := range employeeTestScores {

		totalScore := calculateTotalScore(scores)

		fmt.Printf("===== EMPLOYEE %d PERFORMANCE =====\n", employeeIndex+1)

		for scoreIndex, score := range scores {
			fmt.Printf("Test %d: %.d\n", scoreIndex+1, score)
		}

		averageScore := calculateAverageScore(totalScore, len(scores))

		fmt.Println("Total Score:", totalScore)

		fmt.Printf("Average Score: %.2f\n", averageScore)
	}

}

func main() {

	employeePerformanceTracker()

}
