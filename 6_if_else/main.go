package main

import "fmt"

func main() {
	// age := 16

	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// } else {
	// 	fmt.Println("You are not an adult.")
	// }

	if age := 16; age >= 18 {
		fmt.Println("You are an adult.", age)
	} else if age >= 13 {
		fmt.Println("You are a teenager.", age)
	} else {
		fmt.Println("You are a child.", age)
	}

	// go has no ternary operator like other languages
}
