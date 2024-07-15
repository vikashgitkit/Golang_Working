package main

import "fmt"

func main() {
	var x int = 10
	var p *int

	//p now holds the address of x
	p = &x
	fmt.Println("Values of x is:", x)
	fmt.Println("Memory address of x is:", &x)
	fmt.Println("Value of p (address of x):", p)
	fmt.Println("Value at the address stored in p (value of x):", *p)

	// Modify the value of x using the pointer
	*p = 20
	fmt.Println("New value of x after modification via pointer:", x)

	//Another example code
	y := 15
	fmt.Println("Before y is:", y)
	updateValue(&y)
	fmt.Println("After y is:", y)
}

func updateValue(val *int) {
	//Dereferencing
	*val = 25
	fmt.Println("Val is:", *val)
}
