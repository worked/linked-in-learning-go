package main

import (
	"fmt"
)

// EMPLOYEE
type Employee interface {
	DoWork()
}

type Manager struct {
	Name string
}

type Engineer struct {
	Name string
}

func (m Manager) DoWork() {
	fmt.Printf("Manager %v is doing work\n", m.Name)
}

func (e Engineer) DoWork() {
	fmt.Printf("Engineer %v is doing work\n", e.Name)
}

// PAYROLL
type Payroll struct{}

func (p Payroll) CalculatePayFor(name string) {
	fmt.Printf("Calculating pay for %v\n", name)
}

func main() {
	manager := Manager{Name: "Ed"}
	engineer := Engineer{Name: "Jane"}
	//manager.DoWork("John")
	//engineer.DoWork("Jane")

	employees := []Employee{
		manager,
		engineer,
	}

	payroll := Payroll{}

	for _, employee := range employees {
		employee.DoWork()
		switch employee.(type) {
		case Manager:
			payroll.CalculatePayFor(manager.Name)
		case Engineer:
			payroll.CalculatePayFor(engineer.Name)
		}
	}
}
