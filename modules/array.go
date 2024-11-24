package modules

import "fmt"

func Array() {
	fruits := [4]string{"Apple", "Watermelon", "Grape"} //array
	var langs [3]string
	products := [2]string{"Apple", "Google"}

	langs[0] = "Go"
	langs[1] = "JavaScript"
	langs[2] = "TypeScript"

	fmt.Println(langs[0])
	fmt.Println(langs)
	fmt.Println(products[1])
	fmt.Println(products)

	fmt.Println(fruits)
	fmt.Println(len(fruits))

	fruits[2] = "Kiwi"
	fruits[3] = "Mango"

	fmt.Println(fruits)
	fmt.Println(len(fruits))
}
