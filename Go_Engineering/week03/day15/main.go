package main

import "fmt"

type Worker interface {
	Work()
}

type Employee struct {
	Name string
}

type Contractor struct {
	Name string
}

type Manager struct {
	Name string
}

func (m Manager) Work() {
	fmt.Println(m.Name, "is managing the team")
}

func (c Contractor) Work() {
	fmt.Println(c.Name, "is working")
}

func (e Employee) Work() {
	fmt.Println(e.Name, "is working")
}

func startWork(w Worker) {
	w.Work()
}

func display(value any) {
	switch v := value.(type) {
	case string:
		fmt.Println("This is a string:", v)
	case int:
		fmt.Println("This is an integer:", v)
	case float64:
		fmt.Println("This is a float:", v)
	case bool:
		fmt.Println("This is a boolean:", v)
	case Employee:
		fmt.Println("Employee Name:", v.Name)
	case Contractor:
		fmt.Println("Contractor Name:", v.Name)
	case Manager:
		fmt.Println("Manager Name:", v.Name)
	default:
		fmt.Println("Unknown type")
	}
}

func checkInteger(value any) {
	if num, ok := value.(int); ok {
		fmt.Println("This is an integer:", num)
	} else {
		fmt.Println("This value is not an integer")
	}
}

func main() {
	emp := Employee{Name: "Ali Samuel"}
	con := Contractor{Name: "John Peter"}
	man := Manager{Name: "Samuel"}

	display("Hello Go")
	display(25)
	display(45.5)
	display(true)

	startWork(emp)
	startWork(con)
	startWork(man)

	display(emp)
	display(con)
	display(man)

	checkInteger(100)
	checkInteger("Hello Go")
	checkInteger(45.5)
}
