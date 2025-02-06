package main

func main() {
	age := 34
	if age >= 21 {
		println("You are can drink whisky")
	} else if age <= 18 {
		println("you can drink beer")
	} else {
		println("you can drink water")
	}

	// if statement with a short statement
	age = 21
	if age >= 18 {
		println("you can vote now")

	} else {
		println("you can't vote")
	}

	var role = "admin"
	var hasaccess = true
	if role == "admin" && hasaccess {
		println("you have access")
	} else {
		println("you don't have access")
	}

}
