
package main

import "fmt"

// PACKAGE-LEVEL VARIABLE
var schoolName = "University of Abuja"

func displayStudent() {

	// LOCAL VARIABLES
	studentName := "Kelvin"
	age := 25

	fmt.Println("School:", schoolName)
	fmt.Println("Name:", studentName)
	fmt.Println("Age:", age)

	// BLOCK SCOPE
	if age >= 18 {
		status := "Adult"
		fmt.Println("Status:", status)
	}
}

//EXERCISE ONE
//DECLARING A PACKAGE LEVEL VARIABLE

var companyName = "SimplyTech Solutons"

func displayEmployee() {
	employeeName := "Samuel"
	salary := 150000
	fmt.Printf("EMPLOYEE NAME: %s \n SALARY: %d \n COMPANY NAME: %s \n ", companyName, employeeName, salary)
}

func displayCompany() {
	fmt.Println(companyName)
}

//EXERCISE TWO
//UDERSTANDING BLOCK SCOPE

var age = 22

if age >= 18 {
	status := "Eligible"
	fmt.Println(status)
}

fmt.Println(age)

//EXERCISE THREE
//CREATIING ANONYMOUS FUNCTION THAT ACCEPTS TWO INTEGERS AND ADDS THEEM

func(a int, b int){
	result = a + b
	return result

}

//EXERRCISE FOUR
//IMMEDIATELY INVOKED FUNCTION EXPRESSION(IIFE)

func(){
	fmt.Println("Hey dear! i am anonymous")
}()

func main() {
	displayStudent()
	displayCompany()
	displayEmployee()
}
