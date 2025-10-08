package main

import (
	"fmt"
	"strings"
)

// EMPLOYEE
type Employee interface {
	DoWork(name string)
	GetName() string
}

type Manager struct {
	Name string
}

type Engineer struct {
	Name string
}

func (m Manager) DoWork(name string) {
	fmt.Printf("Manager %v is doing work\n", name)
}

func (e Engineer) DoWork(name string) {
	fmt.Printf("Engineer %v is doing work\n", name)
}

func (m Manager) GetName() string {
	return m.Name
}

func (e Engineer) GetName() string {
	return e.Name
}

// PAYROLL
type Payroll struct{}

func (p Payroll) CalculatePay(employee Employee) {
	employeeType := fmt.Sprintf("%T", employee)
	// Remove package prefix (e.g., "main.Manager" -> "Manager")
	typeName := strings.TrimPrefix(employeeType, "main.")
	fmt.Printf("Calculating pay for %v the %v\n", employee.GetName(), typeName)
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
	for _, employee := range employees {
		employee.DoWork(employee.GetName())
	}

	payroll := Payroll{}
	payroll.CalculatePay(manager)
	payroll.CalculatePay(engineer)
}
