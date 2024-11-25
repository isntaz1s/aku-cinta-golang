package modules

import "fmt"

type User struct {
	Id   int
	Name string
	Age  int
}

type Customer struct {
	Id   int
	Name string
	Addr string
}

type HasName interface {
	getName() string
}

func Struct(user User) User {
	return user
}

func (user User) Greet(name string) {
	fmt.Printf("Hello, %s my name is %s", name, user.Name)
}

func Interface(val any) any {
	return val
}

func Slices(arg []int) {
	if arg == nil {
		fmt.Println("Arg  is nil or not have values")
	}
	fmt.Println(arg)
}

func Maps(arg map[string]string) {
	if arg == nil {
		fmt.Println("Arg is nil or not have values")
	}
	fmt.Println(arg)
}

func NilFunc(f func(arg string), args string) string {
	return args
}

func Nil() {
	user := User{
		Id:   1,
		Name: "Azis",
		Age:  17,
	}

	argSlice := []int{1, 2, 3, 4, 5}
	argMap := map[string]string{
		"Name":  "Car",
		"Brand": "Toyota",
		"Year":  "2022",
	}

	nilFValue := func(arg string) {}

	fmt.Println(NilFunc(nilFValue, "Hi"))

	str := Struct(user)
	if i := Interface(true); i == nil {
		fmt.Println("Interface nil", i)
	} else {
		fmt.Println("Interface not nil")
	}

	fmt.Println(str)
	str.Greet("Budi")
	Slices(argSlice)
	Maps(argMap)
}
