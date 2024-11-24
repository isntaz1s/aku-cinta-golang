package modules

import "fmt"

func Func(arg string) string {
	fmt.Println(arg)
	return arg
}

func VariadicFunc(arg ...int) int {
	box := 0

	for _, value := range arg {
		box += value
	}

	fmt.Println(box)

	return box
}

func FunctionParameter(fc func(arg string) string) {
	fmt.Println(fc("Uhuy"))
}

func FunctionValue(val string) string {
	fmt.Println(val)
	return val
}
