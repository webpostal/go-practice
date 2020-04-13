package main

import "fmt"

func main() {
	x := 1000
	y := 100
	if x < y {
		fmt.Printf("%v is lesser than %v\n", x, y)
	} else if x > y {
		fmt.Printf("%v is greater than %v\n", x, y)
	} else {
		fmt.Printf("%v is equal to %v\n", x, y)
	}

	// Another example

	w := "Marco Polo"

	if w == "Marcopolo" {
		fmt.Println(w)
	} else if w == "Cinderella" {
		fmt.Printf("You got it. It's %v\n", w)
	} else {
		fmt.Println("none of the above!\n")
	}

}
