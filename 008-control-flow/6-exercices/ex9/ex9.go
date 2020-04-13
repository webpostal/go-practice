package main

import "fmt"

func main() {

	favSport := "swimming"

	switch favSport {

	case "hiking":
		fmt.Println("Prepare your hiking tools")
	case "surfing":
		fmt.Println("Surfing is a great sport!")
	case "swimming":
		fmt.Println("Please try not to drown!")

	default:
		fmt.Println("So what's your favorite sport then?")
	}

}
