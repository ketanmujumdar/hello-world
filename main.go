package main

import "fmt"

// calculateSum calculates the sum of a slice of numbers
func calculateSum(numbers []int) int {
    sum := 0
    for _, num := range numbers {
        sum += num
    }
    return sum
}

func main() {
    fmt.Println("Hello, World!")

    // Demonstrate the new feature
    numbers := []int{1, 2, 3, 4, 5}
    sum := calculateSum(numbers)
    fmt.Printf("The sum of numbers %v is: %d\n", numbers, sum)
}