package main

import "fmt"

// Declare a function to calculate the area of a rectangle
func calculateArea(width, height float64) float64 {
    return width * height
}

// Declare a function to get user input
func getUserInput(prompt string) float64 {
    var value float64

    fmt.Print(prompt)
    fmt.Scanf("%f", &value)

    return value
}

func main() {
    // Get user input for width and height
    width := getUserInput("Enter the width of the rectangle: ")
    height := getUserInput("Enter the height of the rectangle: ")

    // Calculate the area
    area := calculateArea(width, height)

    // Print the result
    fmt.Printf("The area of the rectangle is: %.2f\n", area)
}
