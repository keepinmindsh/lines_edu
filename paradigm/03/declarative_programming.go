package main

import (
    "fmt"
)

type predicateFunc func(int) bool

func filter(numbers []int, predicate predicateFunc) []int {
    var result []int
    for _, num := range numbers {
        if predicate(num) {
            result = append(result, num)
        }
    }
    return result
}

func isEven(number int) bool {
    return number%2 == 0
}

func isPositive(number int) bool {
    return number > 0
}

func main() {
    numbers := []int{-3, -2, -1, 0, 1, 2, 3}

    evenNumbers := filter(numbers, isEven)
    fmt.Printf("Even numbers: %v\n", evenNumbers)

    positiveNumbers := filter(numbers, isPositive)
    fmt.Printf("Positive numbers: %v\n", positiveNumbers)
}
