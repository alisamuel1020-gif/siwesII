package main

import "fmt"

func two() {
	fmt.Print("The *Print()*--> This Displays out put on the console ")
	fmt.Print("and remains on thesame line \n")
	fmt.Println("The *Println()*--> this displays a message on the console")
	fmt.Println("John will display John on the Console")
	fmt.Printf("The *Printf* --> this prints formarted strings like")
	name := "Peter"
	age := 30
	state := "Kogi"
	country := "Nigeria"
	fmt.Printf("Hello My name is %s i am %d years old i am from %s state in %s ", name, age, state, country)
}
