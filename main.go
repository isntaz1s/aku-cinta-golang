package main

import (
  "aku-cinta-go/modules"
	"aku-cinta-go/models"
	"fmt"
)

func variables(nm string, ag int, student bool) bool {
	name := nm
	var age int = ag
	var isStudent bool = student

	const (
		isOnline bool = true
		isActive bool = false
	)

	fmt.Println(name, age, isStudent)
	fmt.Println(isOnline)

	return true
}

func alias() string {
	type User string

	type UserStruct struct {
		Id       int
		Username string
	}

	const user User = "User"

	var userStruct UserStruct = UserStruct{Id: 1, Username: "Azis"}

	fmt.Println(userStruct)
	fmt.Printf("%v", userStruct.Id)
	fmt.Printf("%s", userStruct.Username)
	return "Alias"
}

func calc(param ...int) int {

	var num int = 1

	if param != nil {
		fmt.Printf("Err")
	}

	for i := 0; i < len(param); i++ {
		num += param[i]
		fmt.Println(num)
	}

	return 1
}

func students(student []string) []string {

	box := ""

	for i := 0; i < len(student); i++ {
		box += student[i]
		fmt.Println(box)
	}
	return student
}

func Users(data models.User) models.User {
	var user1 models.User = models.User{Id: data.Id, Username: data.Username}
	user := models.UserWithEmail{Id: 2, Username: "Azis", Email: "azis@yahoo.co.id"}
	box := []string{}

	box = append(box, user1.Username)
	box = append(box, user.Email)

	fmt.Println(user1)
	fmt.Println(user)
	fmt.Println(box)
	return data
}

func main() {
	vrb := variables("Azis", 17, true)
	als := alias()
	nums := calc(1, 2, 3, 4, 5)

	student := []string{"Azis", "Abdul"}
	stds := students(student)

	data := models.User{Id: 1, Username: "Abdul"}
	users := Users(data)

	fmt.Println(vrb)
	fmt.Println(als)
	fmt.Println(nums)
	fmt.Println(stds)
	fmt.Println(users)
  modules.Hello("Azis")
  modules.Calculate(1,2,3,4,5)
}
