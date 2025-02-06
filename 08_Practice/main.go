package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var age int = 25
	const pi float32 = 3.14
	var name string = "Bittu Singh"
	println(age, pi, name)

	// var i = 0
	// for i=0; i<=100; i++{
	// 	println(i)
	// }

	var num = 7
	if num%2 == 0 {
		println("Even")
	} else {
		println("odd")
	}

	switch 3 {
	case 1:
		println("Monday")
	case 2:
		println("Tuesday")
	case 3:
		println("Wednesday")
	case 4:
		println("Thursday")
	case 5:
		println("Friday")
	case 6:
		println("Saturday")
	case 7:
		println("Sunday")
	default:
		println("Invalid")
	}

	num1, num2, operator := 5, 3, "+"
	switch operator {
	case "+":
		println("Addition is", num1+num2)
	case "-":
		println("Subtraction is", num1-num2)
	case "*":
		println("Multiplication is", num1*num2)
	case "/":
		println("Division is", num1/num2)
	default:
		println("Invalid Operator")
	}
}
