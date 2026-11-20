package main

import "fmt"

func arraysAndSlices() {
	employeeAges := make([]int, 3, 5)

	employeeAges[0] = 23
	employeeAges[1] = 25
	employeeAges[2] = 28

	fmt.Println("Initial slice:", employeeAges)

	employeeAges = append(employeeAges, 30)

	fmt.Println("Updated slice:", employeeAges)

	fmt.Println("Length:", len(employeeAges))
	fmt.Println("Capacity:", cap(employeeAges))
}

func displayEmployeeAges(ages []int) {

	for index, age := range ages {
		fmt.Printf("Index: %d Age: %d\n", index, age)
	}
}

func updateAge(ages []int) []int {
	ages = append(ages, 40)
	return ages
}

func main() {

	employeeAges := []int{23, 25, 28, 30, 32}

	employeeAges = updateAge(employeeAges)

	arraysAndSlices()

	displayEmployeeAges(employeeAges)

	fmt.Println("Updated slice:", employeeAges)
}
