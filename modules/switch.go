package modules

import "fmt"

func Switch() {
  fruits := "Apple"

  switch fruits {
    case "Apple":
      fmt.Println("It is an apple!")
    case "Mango":
      fmt.Println("It is a mango!")
    default:
      fmt.Println("It is not a fruit!")
  }

  switch len(fruits) <= 5 {
    case true:
      fmt.Println("It is a fruit!")
    case false:
      fmt.Println("It is not a fruit!")
  }

  switch {
    case fruits == "Apple":
      fmt.Println("It is an apple!")
    default:
      fmt.Println("It is not a fruit!")
  }
}
