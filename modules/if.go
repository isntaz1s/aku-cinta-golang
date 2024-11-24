package modules

import "fmt"

func If() {
	name := "Azis"

	if name == "Azis" {
		fmt.Println("Hello Azis")
	} else if name == "Abdul" {
		fmt.Println("Hello Abdul")
	} else {
		fmt.Println("Hi, what is your name?")
	}

	if len(name) <= 5 {
		fmt.Println("Right name!")
	} else if len(name) <= 10 {
		fmt.Println("Also, right name!")
	} else {
		fmt.Println("Name is too long!")
	}
}
