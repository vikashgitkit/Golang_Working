package main

import "fmt"

func concat(strs ...string) string {
	final := ""

	//strs is just a slice of strings
	for i := 0; i < len(strs); i++ {
		final += strs[i]
		fmt.Println("Final value is:", final)
	}
	fmt.Println("Final value before return is:", final)
	return final
}

func main() {
	final := concat("Hello", "there", "friend")
	fmt.Println("Final inside main fn:", final)
}
