package main

import "fmt"

func main() {

    fmt.Println("Welcome to my quiz game. Let's start!")

    // Ask for user input for a name
    var name string
    fmt.Printf("Enter your name: ") 
    fmt.Scan(&name)

    // Print the greeting
    fmt.Printf("Hello %v, welcome to the game!\n", name)

    // Ask for user input for an age
    var age uint
    fmt.Printf("Enter your age (in number format): ") 
    fmt.Scan(&age)

    // Print with condition
    if age >= 10 {
        fmt.Println("Yay, you can play!")
    } else {
        fmt.Println("Sorry, you are too young to play!")
        return // Exit the game if age is less than 10
    }

    // Define score
    var score = 0
    var num_questions = 2

    // Ask for user input about graphics card
    var answer string
    var answer2 string
    fmt.Printf("What is better, the RTX 3080 or RTX 3090? Answer with two words! \n")
    fmt.Scan(&answer, &answer2)

    if answer + " " + answer2 == "RTX 3090" || answer + " " + answer2 == "rtx 3090" {
        fmt.Println("Correct!")
        score++
    } else {
        fmt.Println("Incorrect!")
    }

    // Ask for user input about CPU
    var cores int
    fmt.Printf("How many cores does the Ryzen 9 3900x have? In number format \n")
    fmt.Scan(&cores)

    if cores == 12 {
        fmt.Println("Correct!")
        score++
    } else {
        fmt.Println("Incorrect!")
    }

    // Output score and percentage
    fmt.Printf("You scored %v out of %v.\n", score, num_questions)
    
    // Calculate percentage
    percent := (float64(score) / float64(num_questions)) * 100
    fmt.Printf("You scored: %.2f%%.\n", percent) // Use %.2f to format percentage to 2 decimal places
}
