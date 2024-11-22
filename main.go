package main

import (
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
    Id int
    Username string
  }

  const user User = "User"

  var userStruct UserStruct = UserStruct{ Id: 1, Username: "Azis"}

  fmt.Println(userStruct)
  fmt.Printf("%v", userStruct.Id)
  fmt.Printf("%s", userStruct.Username)
  return "Alias"
}

func calc(param ...int) (int) {

  var num int = 1

  if param != nil {
    fmt.Printf("Err")
  }

  for i :=0; i < len(param); i++ {
    num += param[i]
    fmt.Println(num)
  }

  return 1
}

func students(student []string) ([]string) {

  box := ""

  for i := 0; i < len(student); i++ {
    box += student[i]
    fmt.Println(box)
  }
  return student
}

func main() {
  vrb := variables("Azis", 17, true)
  als := alias()
  nums := calc(1,2,3,4,5)

  student := []string{"Azis", "Abdul"}
  stds := students(student)

  fmt.Println(vrb)
  fmt.Println(als)
  fmt.Println(nums)
  fmt.Println(stds)
}
