package main

import (
	"fmt"
	"time"
)

func SwitchWithNoCondtion() {

	t := time.Now()
	fmt.Println("Data is t:", t)

	// The switch statement is used without an explicit condition. This means it will evaluate each case as a boolean expression.
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning..")
	case t.Hour() < 17:
		fmt.Println("Good afternoon..")
	default:
		fmt.Println("Good evening..")
	}
}
