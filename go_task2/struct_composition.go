package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func main() {
	emp := Employee{
		Person:     Person{Name: "Alice", Age: 30},
		EmployeeID: "E12345",
	}

	fmt.Printf("Employee Name: %s, Age: %d, Employee ID: %s\n", emp.Name, emp.Age, emp.EmployeeID)
}
