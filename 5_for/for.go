package main

import "fmt"

// For loop in Go is used to iterate over a range of values, similar to for loops in other programming languages.
// only construct in go for looping
func main() {
	// while loop
	i := 0
	for i <= 2 {
		fmt.Println("Hello World", i)
		i++
	}
	// infinite loop
	// for {
	// 	fmt.Println("Infinite Loop")
	// 	// break // Uncomment to stop the infinite loop
	// }

	// classic for loop
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	// for range loop
	for i := range "Golang" {
		fmt.Println(i)
	}

}
