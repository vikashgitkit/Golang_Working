package main

import "fmt"

func main() {
	//Case1
	// primes := [6]int{2, 3, 4, 5, 6, 7}

	// var s []int = primes[1:4]
	// fmt.Println(s)

	//Case2
	// names := [4]string{
	// 	"John",
	// 	"Paul",
	// 	"George",
	// 	"Ringo",
	// }

	// fmt.Println("Names are:", names)

	// a := names[0:2]
	// b := names[1:3]
	// fmt.Println("Here are names in A and B:", a, b)

	// b[0] = "AAA"
	// fmt.Println("Updated A and B is:", a, b)
	// fmt.Println("Last Names are:", names)

	//Case3- Slice literals
	// q := []int{2, 3, 5, 6, 7, 8, 99, 76}
	// fmt.Println("Q is:", q)

	// r := []bool{true, false, true, true, false, true}
	// fmt.Println(r)

	// s := []struct {
	// 	i int
	// 	b bool
	// }{
	// 	{2, true},
	// 	{3, false},
	// 	{5, true},
	// 	{7, true},
	// 	{11, false},
	// 	{13, true},
	// }
	// fmt.Println("S is:", s)

	//Case4- Slices defaults
	// s := []int{2, 3, 5, 7, 11, 13}

	// s = s[1:4]
	// fmt.Println(s)

	// s = s[:2]
	// fmt.Println(s)

	// s = s[1:]
	// fmt.Println(s)

	//Case5- Slice Length and capacity
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
