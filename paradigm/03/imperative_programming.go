package main 

import (
	"fmt"
)

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result = result * i
	}
	return result
}

func main(){
  var number int
	fmt.Print("Enter a number to calculate its factorial: ")
	fmt.Scan(&number)

	fact := factorial(number)
	fmt.Printf("Factorial of %d is: %d\n", number, fact)
}
