package main

import "fmt"

func main() {
    var number int
    fmt.Print("Enter an integer: ")
    fmt.Scanln(&number)

    if number > 0 {
        fmt.Println("The number is positive.")
    } else if number < 0 {
        fmt.Println("The number is negative.")
    } else {
        fmt.Println("The number is zero.")
    }
}


// func main() {
//     sum := 0
//     for i := 1; i <= 10; i++ {
//         sum += i
//     }
//     fmt.Println("The sum of the first 10 natural numbers is:", sum)
// }


// func main() {
//     var day int
//     fmt.Print("Enter a number (1 for Monday, 7 for Sunday): ")
//     fmt.Scanln(&day)

//     switch day {
//     case 1:
//         fmt.Println("Monday")
//     case 2:
//         fmt.Println("Tuesday")
//     case 3:
//         fmt.Println("Wednesday")
//     case 4:
//         fmt.Println("Thursday")
//     case 5:
//         fmt.Println("Friday")
//     case 6:
//         fmt.Println("Saturday")
//     case 7:
//         fmt.Println("Sunday")
//     default:
//         fmt.Println("Invalid day number.")
//     }
// }
// package main

// import "fmt"

// func days() {
//     var day int
//     daysOfWeek := [7]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

//     fmt.Print("Enter a number (1 for Monday, 7 for Sunday): ")
//     fmt.Scanln(&day)

//     if day >= 1 && day <= 7 {
//         fmt.Println(daysOfWeek[day-1])
//     } else {
//         fmt.Println("Invalid day number.")
//     }
// }
