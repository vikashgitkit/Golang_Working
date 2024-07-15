package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go Runs On: ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")

	default:
		fmt.Printf("%s.\n", os)
	}

	Main1()
	SwitchWithNoCondtion()

}
