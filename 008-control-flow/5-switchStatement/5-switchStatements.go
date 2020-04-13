package main

import "fmt"

func main() {

	switch {

	case (2 == 2):
		fmt.Println("Eval to true")
	case (2 == 6):
		fmt.Println(false)

	}

	switch "Rossina" {

	case "Camila":
		fmt.Println("Camila will not Print")
	case "Georgina":
		fmt.Println("Georgina will not Print")
	case "Rossina":
		fmt.Println("Rossina will Print")
	default:
		fmt.Println("Default value")
	}

	x := "River"
	switch x {

	case "Mountain":
		fmt.Println("Camila will not Print")
	case "River":
		fmt.Println("Georgina will not Print")
	case "Lake":
		fmt.Println("Rossina will Print")
	default:
		fmt.Println("Default value")
	}

}
