package modules

import "fmt"

func Calculate(arg ...int) []int {
	if arg != nil {
		fmt.Println("Args required")
	}

	sum := 0

	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}

	fmt.Println("The sum is", sum)

	return arg
}
