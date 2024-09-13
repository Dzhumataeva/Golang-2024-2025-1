package main

import "fmt"

func main() {
    // Using var keyword
    var age int = 25
    var height float64 = 5.9
    var name string = "Alice"
    var isStudent bool = true

    // Using short declaration syntax
    weight := 68.5
    country := "USA"
    passed := false

    // Printing values and types using fmt.Printf
    fmt.Printf("Age: %d, Type: %T\n", age, age)
    fmt.Printf("Height: %.1f, Type: %T\n", height, height)
    fmt.Printf("Name: %s, Type: %T\n", name, name)
    fmt.Printf("Is Student: %t, Type: %T\n", isStudent, isStudent)
    fmt.Printf("Weight: %.1f, Type: %T\n", weight, weight)
    fmt.Printf("Country: %s, Type: %T\n", country, country)
    fmt.Printf("Passed: %t, Type: %T\n", passed, passed)
}
