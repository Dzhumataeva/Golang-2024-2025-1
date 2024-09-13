package main

import "fmt"

type Employee struct {
    Name string
    ID   int
}

type Manager struct {
    Employee
    Department string
}

//  создатб метод для  Employee
func (e Employee) Work() {
    fmt.Printf("Employee %s with ID %d is working.\n", e.Name, e.ID)
}

// Создать экземпляр Manager и вызов метод Work
func main() {
    manager := Manager{
        Employee:   Employee{Name: "John", ID: 101},
        Department: "Sales",
    }

    manager.Work() 
}
