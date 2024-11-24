package modules

import "fmt"

type Arg string

func Hello(arg Arg) {
	fmt.Printf("Hi %s, welcome!", arg)
}
