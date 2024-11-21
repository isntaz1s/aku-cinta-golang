package main

import "fmt"

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

func main() {
  vrb := variables("Azis", 17, true)

  fmt.Println(vrb)
}
