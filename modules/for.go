package modules

import "fmt"

func For() {
  counter := 0

  brandArray := [3]string{"Mitsubishi", "Yamaha", "Honda"}
  brandSlice := []string{"Mitsubishi", "Yamaha", "Honda"}
  brandMap := map[string]string{
    "car": "Mitsubishi",
    "motorcycle": "Yamaha",
  }

  for counter <= 10 {
    fmt.Println("Counter = ", counter)
    counter++
  }

  for count := 0; count <= 10; count++ {
    fmt.Println("Count = ", count)
  }

  for index, value := range brandArray{
    fmt.Println("Index: ", index, "Value: ", value)
  }

  for _, value := range brandSlice {
    fmt.Println("Value: ", value)
  }

  for _, value := range brandMap {
    fmt.Println("Value: ", value)
  }
}
