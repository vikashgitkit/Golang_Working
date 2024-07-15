package main

import (
	"fmt"
)

// Define a struct type
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Implement the Stringer interface for the Person type
func (p Person) String() string {
	return fmt.Sprintf("%s %s (Age: %d)", p.FirstName, p.LastName, p.Age)
}

func main() {
	// Create an instance of Person
	p := Person{FirstName: "John", LastName: "Doe", Age: 30}

	// Print the instance using fmt package
	fmt.Println(p)
}
