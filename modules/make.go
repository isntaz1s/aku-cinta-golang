package modules

import "fmt"

func Make() {
	slice := make([]string, 4)

	slice[0] = "Hello"
	slice[1] = "World!"

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(append(slice, "Grateful"))
}
