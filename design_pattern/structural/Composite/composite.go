package main

import "fmt"

// Component interface
type Employee interface {
	GetDetails() string
}

// Leaf: Developer
type Developer struct {
	name   string
	salary int
}

func (d *Developer) GetDetails() string {
	return fmt.Sprintf("Developer: %s, Salary: %d", d.name, d.salary)
}

// Composite: Manager
type Manager struct {
	name         string
	subordinates []Employee
}

func (m *Manager) Add(employee Employee) {
	m.subordinates = append(m.subordinates, employee)
}

func (m *Manager) GetDetails() string {
	details := fmt.Sprintf("Manager: %s\n", m.name)
	for _, employee := range m.subordinates {
		details += "  " + employee.GetDetails() + "\n"
	}
	return details
}

func main() {
	dev1 := &Developer{name: "Alice", salary: 80000}
	dev2 := &Developer{name: "Bob", salary: 90000}

	manager := &Manager{name: "Eve"}
	manager.Add(dev1)
	manager.Add(dev2)

	fmt.Println(manager.GetDetails())
}
