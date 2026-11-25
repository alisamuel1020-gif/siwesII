package main

import "fmt"

var age int = 25
var agePointer *int

func nilPointer() {
	if agePointer != nil {
		fmt.Println(*agePointer)
	} else {
		fmt.Println("Pointer is nil. Assigning an address...")
		agePointer = &age
	}

	fmt.Println("Age:", *agePointer)

}
