package main

import "fmt"

// Function to add two integers
func add(a int, b int) int {
    return a + b
}

func main() {
    sum := add(3, 5)
    fmt.Println("Sum:", sum)
}


// Function to swap two strings
// func swap(x, y string) (string, string) {
//     return y, x
// }

// func main() {
//     a, b := swap("Hello", "World")
//     fmt.Println("Swapped:", a, b)
// }


// func divide(a int, b int) (int, int) {
//     return a / b, a % b
// }

// func main() {
//     quotient, remainder := divide(10, 3)
//     fmt.Println("Quotient:", quotient, "Remainder:", remainder)
// }
