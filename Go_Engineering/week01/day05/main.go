package main

import "fmt"

func main() {
	//EXERCISE ONE: NUMCER PROCESSOR
	//1 PRINTING 1 TO 20
	for i := 1; i <= 20; i++ {
		fmt.Println(i)
	}

	//2 PRINTING EVEN NUMBERS FROM 1 TO 20
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	//3 PRINTING NUMBERS FROM 1 TO 20 BUT SKIPPING 10.
	for i := 1; i <= 20; i++ {
		if i == 10 {
			continue
		}
		fmt.Println(i)
	}

	//4 PRINTING NUMBERS FROM 1 TO 20 BUT STOPPING AT 15.
	for i := 1; i <= 20; i++ {
		if i == 15 {
			break
		}
		fmt.Println(i)
	}

	//EXERCISE TWO: MULTIPLICATION TABLE
	//1 PRINTING MULTIPLICATION TABLE OF 5
	for i := 1; i <= 10; i++ {
		fmt.Printf("5 x %d = %d\n", i, 5*i)
	}

	//EXERCISE THREE: STUDENT SCORE PROCESSOR
	scores := [5]int{70, 65, 80, 75, 90}

	//Displaying all five scores using a loop
	for i := 0; i < 5; i++ {
		fmt.Printf("Score %d: %d\n", i+1, scores[i])
	}

	//Calculating total score
	totalScore := 0
	for i := 0; i < 5; i++ {
		totalScore += scores[i]
	}

	//Calculating average score
	averageScore := float64(totalScore) / 5
	fmt.Printf("Total Score: %d\n", totalScore)
	fmt.Printf("Average Score: %.2f\n", averageScore)

	//FiINDING THE HIGHEST SCORE
	highestScore := scores[0]
	for i := 1; i < 5; i++ {
		if scores[i] > highestScore {
			highestScore = scores[i]
		}
	}
	fmt.Printf("Highest Score: %d\n", highestScore)

	//Finding the lowest score
	lowestScore := scores[0]
	for i := 1; i < 5; i++ {
		if scores[i] < lowestScore {
			lowestScore = scores[i]
		}
	}
	fmt.Printf("Lowest Score: %d\n", lowestScore)

	//COUNTING How MANY STUDENTS' SCORES ARE 70 OR ABOVE
	count := 0
	for i := 0; i < 5; i++ {
		if scores[i] >= 70 {
			count++
		}
	}
	fmt.Printf("Number of students with scores 70 or above: %d\n", count)

}
