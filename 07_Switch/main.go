package main

import (
	"fmt"
	"time"
)

func main() {

	// Switch statement
	// Switch statement is used to compare the value of a variable to multiple values.
	// The switch statement is a multi-way branch statement. It provides an easy way to dispatch

	// switch 1 {
	// case 1:
	// 	println("one")

	// case 2:
	// 	println("two")

	// case 3:
	// 	println("three")

	// default:
	// 	println("other")

	// }

	//we can also use mutliple  swicth in go
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is weekend")
	default:
		fmt.Println("It is weekday")

	}

	whoamI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			println("I am an integer")
		case string:
			fmt.Println("I am a string", t)
		case bool:
			println("I am a boolean")

		default:
			println("I am another type")
		}
	}
	whoamI(21.5)
}
