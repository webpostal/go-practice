package main

import (
	"fmt"
)

func main() {

	// In this first example we limit the scope of x to the first if statement

	if x := 42; x == 42 {
		fmt.Println("001")
	}

	y := 43
	if y == 43 {
		fmt.Println("002")
	}
	fmt.Printf("%v - here 'y' is scoped to the whole main func\n", y)

	// Go will complety the semicolon at the end of each line, but
	// we can get two statements in on line using ;
	fmt.Println("First one");fmt.Println("Second one")

}
