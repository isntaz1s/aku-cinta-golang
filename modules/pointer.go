package modules

import "fmt"

type Shape struct {
	Type string
}

type Person struct {
	Id   int
	Name string
}

type Address interface {
	HasCity() string
}

func (person *Person) GetName() {
	fmt.Println("Name", person.Name)
}

func (person *Person) HasCity() {
	fmt.Println("City", person)
}

func Pointer() {
	var block *Shape = &Shape{"Block"}
	var circle *Shape = &Shape{"Circle"}

	num := 0
	num2 := &num

	*num2 = 5

	fmt.Println(num)
	fmt.Println(num2)

	person := Person{
		Id:   1,
		Name: "Azis",
	}

	person.GetName()

	person.Name = "Abdul"

	person.GetName()
	person.HasCity()

	fmt.Println(block)
	fmt.Println(circle)
}
