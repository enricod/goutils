package main

import (
	"fmt"

	"github.com/enricod/gotools/fp"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Map: square each number
	squares := fp.Map(numbers, func(n int) int {
		return n * n
	})
	fmt.Println("Squares:", squares)

	// Filter: keep only even numbers
	evens := fp.Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Evens:", evens)

	// Reduce: sum of numbers
	sum := fp.Reduce(numbers, 0, func(acc, n int) int {
		return acc + n
	})
	fmt.Println("Sum:", sum)

	// Reduce: concatenate as strings
	joined := fp.Reduce(numbers, "", func(acc string, n int) string {
		return acc + fmt.Sprintf("%d,", n)
	})
	fmt.Println("Joined:", joined)
}
