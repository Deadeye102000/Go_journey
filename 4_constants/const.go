package main

import "fmt"

const age = 22

// My_name := "Anany" // short variable declaration
// Note: Constants in Go are declared using the `const` keyword and cannot be reassigned.

func main() {
	const name string = "Golang" // explicit type constant

	fmt.Println(name)

	// name = "javascript" // This will cause an error because constants cannot be reassigned

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println("Server running on", host, ":", port)
}
