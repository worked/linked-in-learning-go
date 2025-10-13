package main

import (
	"fmt"
)

// TYPE ASSERTION

type Employee interface {
	DoWork()
}

type Contractor interface {
	DoWork()
}

type Manager struct {
	Name string
}

type Engineer struct {
	Name string
}

type TempEngineer struct {
	Name string
}

func main() {
	// i can be any type, in this case it is a string
	var i interface{} = "hello"

	// idiomatic way to check if the type is correct
	// in this case i is a string "hello"
	if s, ok := i.(string); ok {
		fmt.Println("i is a string:", s)
	} else {
		fmt.Println("i is not a string")
	}

	var manager = Manager{Name: "Alice"}
	manager.DoWork()

	var engineer = Engineer{Name: "Bob"}
	engineer.DoWork()

	var contractor = TempEngineer{Name: "Charlie"}
	contractor.DoWork()

	/*
		// more readable for me, casting the Manager type to Employee interface
		var manager Employee = Manager{Name: "Alice"}
		manager.DoWork()
			// type assertion - checks if Manager is type of Employee
			m, ok := manager.(Manager)
			if ok {
				fmt.Println("Manager:", m.Name)
			} else {
				fmt.Println("Not a Manager")
			}
	*/
}

func (m Manager) DoWork() {
	fmt.Printf("Manager %v is doing work\n", m.Name)
}

func (e Engineer) DoWork() {
	fmt.Printf("Engineer %v is doing work\n", e.Name)
}

func (t TempEngineer) DoWork() {
	fmt.Printf("Contractor %v is doing work\n", t.Name)
}
