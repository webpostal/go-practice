package main

import "fmt"

func main() {
	x := 101.0
	y := 9.0

	if x < y {
		fmt.Printf("%v is lesser than %v\n", x, y)
	}
	if x > y {
		fmt.Printf("%v is greater than %v\n", x, y)
	}

	if x == y {
		fmt.Printf("%v is equal to %v\n", x, y)
	}

}
