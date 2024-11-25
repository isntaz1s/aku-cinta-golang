package modules

import "fmt"

func logging() {
	if r := recover(); r != nil {
		fmt.Println("Error occured: ", r)
	}
	fmt.Println("logging")
}

func Panic(arg string) {
	defer logging()
	if arg == "" {
		panic("Arg must be at least a char of string")
	}
	fmt.Println(arg)
}
