package main

import "fmt"

func main() {
	//Switch statement with no switch expression specified

	switch {
	case (2 == 1):
		fmt.Println("Should not print")
	case false:
		fmt.Println("Should not print erither!")
	case true:
		fmt.Println("This one should print")
		fallthrough

	default:

		fmt.Println("This default value should print because of the fallthrough statement ")
	}

}
