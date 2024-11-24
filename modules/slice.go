package modules

import "fmt"

func Slice() {
	person := [...]string{"Name", "Address", "City"}
	month := [...]string{"Jan", "Feb", "Mar", "Apr", "May"}
	days := make([]string, 7)

	newPerson := person[2:]

	newPerson2 := append(newPerson, "Country")

	newMonth := month[0:2]
	newDays := append(days, "Sunday", "Monday")

	insertMonth := append(newMonth, "Jun")

	fmt.Println(person)
	fmt.Println(newPerson)
	fmt.Println(newPerson2)
	fmt.Println(len(newPerson2))
	fmt.Println(insertMonth)
	fmt.Println(newDays)
}
