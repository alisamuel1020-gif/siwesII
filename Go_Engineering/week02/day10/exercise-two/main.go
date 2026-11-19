package main

import "fmt"

// PART A

var department = "Information Technology"

func displayDepartment() {
	department := "Human Resources"
	fmt.Println(department)

	//PART C
	if true {
		message := "Welcome to the HR Department"
		fmt.Println(message)
	}
	//fmt.Println(message)

}

// PART B

func displayOriginalDepartment() {
	fmt.Println(department)
}

// PART D

func main() {
	displayDepartment()
	displayOriginalDepartment()
}
