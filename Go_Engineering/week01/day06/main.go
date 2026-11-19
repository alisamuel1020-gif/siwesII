package main

import "fmt"

func main() {
	// A DYNAMIC STUDENT RECORD MANAGER USING SLICES
	//1 CREATING AN INITIAL EMPTY SLICE OF STUDENTS SCORES
	scores := make([]int, 0)

	// ADDING SCORES USING APPEND
	scores = append(scores, 70)
	scores = append(scores, 65)
	scores = append(scores, 80)
	scores = append(scores, 75)
	scores = append(scores, 90)

	// DISPLAYING ALL SCORES
	fmt.Println(scores)

	// ADDING MORE SCORES USING APPEND
	scores = append(scores, 85, 95)

	// DISPLAYING UPDAATED SCORES
	fmt.Println(scores)

	//CALCULATING THE TOTAL SCORE
	total := 0
	for _, score := range scores {
		total += score
	}
	fmt.Println("Total:", total)

	// CALCULATING THE AVERAGE SCORE
	average := float64(total) / float64(len(scores))
	fmt.Println("Average:", average)

	//FINDING THE HIGHEST SCORE
	highest := scores[0]
	for _, score := range scores {
		if score > highest {
			highest = score
		}
	}

	// FINDING THE LOWEST SCORE
	lowest := scores[0]
	for _, score := range scores {
		if score < lowest {
			lowest = score
		}
	}

	fmt.Println("Highest Score:", highest)
	fmt.Println("Lowest Score:", lowest)

	//COUNTING HOW MANY STUDENTS SCORED 70 0R ABOVE.
	count := 0
	for _, score := range scores {
		if score >= 70 {
			count++
		}
	}
	fmt.Println("Students scoring 70 or above:", count)

	// REMOVING SCORE AT INDEX 2
	scores = append(scores[:2], scores[3:]...)

	// DISPLAYING THE FINAL SCORES.
	fmt.Println("Final Scores:", scores)

	// DIPLAYING THE FINAL SLICE LENGTH AND CAPACITY
	fmt.Println("Length:", len(scores))
	fmt.Println("Capacity:", cap(scores))
}
