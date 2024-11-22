package modules

import "fmt"

func Map() {
  cars := map[string]string{
    "brand": "Mitsubishi",
    "details": "New Mitsubishi car brand",
  }

  books := make(map[string]string)

  books["title"] = "Go programming language book"
  books["desc"] = "Go programming language"
  books["author"] = "..."

  fmt.Println(cars)
  fmt.Println(cars["brand"])
  fmt.Println(cars["details"])
  fmt.Println(books)
  fmt.Println(books["title"])
}
