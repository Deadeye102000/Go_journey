package main

import (
	"fmt"
	"time"
)

func main() {
	// switch statement in Go is used to execute one of many possible blocks of code based on the value of a variable or expression.
	// It is similar to switch statements in other programming languages.

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Tuesday":
		fmt.Println("Second day of the week")
	case "Wednesday":
		fmt.Println("Midweek")
	case "Thursday":
		fmt.Println("Almost weekend")
	case "Friday":
		fmt.Println("End of the work week")
	default:
		fmt.Println("It's the weekend!")
	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")

	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("I am an int")
		case string:
			fmt.Println("I am a string")
		case bool:
			fmt.Println("I am a bool")
		default:
			fmt.Println("I am of a different type", t)
		}
	}

	whoAmI(42)
	whoAmI("Hello")
	whoAmI(true)
	whoAmI(3.14)
}
