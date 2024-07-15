package main

import "fmt"

func main() {
	m := make(map[string]int)

	//Intializing
	m["Ans"] = 42
	fmt.Println("The Value:", m["Ans"])

	//Updating the mapping
	m["Ans"] = 54
	fmt.Println("Updated value:", m["Ans"])

	//Delete from the map
	delete(m, "Ans")
	fmt.Println("Value After deleting:", m["Ans"])

	v, ok := m["Ans"]
	fmt.Println("The value:", v, "Present?", ok)

}
