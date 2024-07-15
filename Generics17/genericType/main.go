package main

import "fmt"

// Define a generic type with two type parameters
type Pair[T, U any] struct {
	First  T
	Second U
}

func main() {
	intPair := Pair[int, int]{First: 1, Second: 2}
	stringPair := Pair[string, string]{First: "Hello", Second: "World"}
	mixedPair := Pair[string, int]{First: "age", Second: 30}

	fmt.Println(intPair)    // Output: {1 2}
	fmt.Println(stringPair) // Output: {hello world}
	fmt.Println(mixedPair)  // Output: {age 30}
}
