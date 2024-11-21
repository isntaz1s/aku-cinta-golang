package main

import "fmt"

func variables(nm string, ag int, student bool) bool {
  name := nm
  var age int = ag
  var isStudent bool = student

  const (
    firstName string = "Abdul"
    lastName = "Azis"
  )

  fmt.Println(name, age, isStudent)
  fmt.Println(firstName, lastName)

  return true
}

func main() {
  vrb := variables("Azis", 17, true)

  fmt.Println(vrb)
}
