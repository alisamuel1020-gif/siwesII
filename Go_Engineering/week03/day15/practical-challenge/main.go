package main

import (
	"fmt"
)

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
	fmt.Println(e.Name, "is performing assigned tasks")
}

func assignWork(w Worker) {
	w.Work()
}

func displayWorkerType(value any) {

	switch val := value.(type) {
	case Employee:
		fmt.Println("Employee Name:", val.Name)
	case Contractor:
		fmt.Println("Contractor Name:", val.Name)
	case Manager:
		fmt.Println("Manager Name:", val.Name)
	default:
		fmt.Println("Unknown type")
	}

}

func checkWorker(value any) {
	if _, ok := value.(Worker); ok {
		fmt.Println("This worker can perform tasks")
	} else {
		fmt.Println("This value is not a worker")
	}
}

func main() {
	emp := Employee{
		Name: "Ali Samuel",
	}

	con := Contractor{
		Name: "John Peter",
	}

	mager := Manager{
		Name: "Samuel",
	}
	assignWork(emp)
	assignWork(con)
	assignWork(mager)

	displayWorkerType(emp)
	displayWorkerType(con)
	displayWorkerType(mager)

	checkWorker(emp)
	checkWorker(con)
	checkWorker(mager)
	checkWorker("Hello Go")
	checkWorker(100)
}
