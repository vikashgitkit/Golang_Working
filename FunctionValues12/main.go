package main

import "fmt"

// Passing function as an arg
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main() {
	//Annonymous fn
	add := func(a, b int) int {
		return a + b
	}

	result := applyOperation(3, 4, add)
	fmt.Println(result) // Output: 7
}
