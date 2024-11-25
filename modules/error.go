package modules

import (
	"errors"
	"fmt"
)

func sayHello(name string) (string, error) {
	if name == "Azis" {
		fmt.Println("Hello!")
	} else {
		fmt.Println("Hi")
	}

	return name, nil
}

func Error() {
	var (
		errValidation = errors.New("Error Validation")
		errUnknown    = errors.New("Error Unknown")
	)

	name, err := sayHello("Azis")

	if err != nil {
		fmt.Println(errValidation)
	} else {
		fmt.Println(errUnknown)
	}

	fmt.Println(name)
}
