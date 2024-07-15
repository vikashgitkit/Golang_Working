package main

import (
	"fmt"
	"os"
)

func main() {
	// Open a file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Ensure the file is closed after the function completes
	defer file.Close()

	// Read from the file
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Read %d bytes: %s\n", count, data[:count])
}
