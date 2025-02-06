package main

import "fmt"

const name = "hello" //we can also declare out of the fuction
// name:= "Bittu"this will give an error because we can't change the value of constant and also we can' declare it using :=
func main() {
	// const name = "Bittu"//constant in go means a variable that can't be changed
	fmt.Println(name)

	const (
		name = "Bittu"
		age  = 21
	)

	fmt.Println(name, age)
}
