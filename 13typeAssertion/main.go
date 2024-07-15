package main

import (
	"fmt"
)

// Define an interface
type Animal interface {
	Speak() string
}

// Define a struct type
type Dog struct {
	Name string
}

// Implement the interface for dag
func (d Dog) Speak() string {
	return "Woof!"
}

// Define another struct type
type Cat struct {
	Name string
}

// Implement the interface for cat
func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	//Create a slice of Animal interface
	animals := []Animal{
		Dog{Name: "Buddy"},
		Cat{Name: "Missy"},
	}

	for _, animal := range animals {
		fmt.Println("Animal is:", animal)
		//Single value assertion(unsafe, will panic if it fails)
		// dog := animal.(Dog)
		// fmt.Println("Dog's is:", dog)
		// fmt.Println("Dog's name is:", dog.Name)

		//Comma-ok(ok is boolean) idiom(Safe way to handle type assertions)
		if dog, ok := animal.(Dog); ok {
			fmt.Println("Dog's name:", dog.Name)

		} else if cat, ok := animal.(Cat); ok {
			fmt.Println("Cat's name:", cat.Name)
		} else {
			fmt.Println("Unknown animal!")
		}
	}
}
