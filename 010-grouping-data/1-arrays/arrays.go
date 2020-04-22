package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)

	x[1] = 9
	fmt.Println(x)

	x[4] = 14
	fmt.Println(x)

	// Composite Literal: [3]string{a, b, c}
	y := [4]string{"One", "Two", "Three", "Four"}
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%q\n", y)

	y[3] = "Five"
	fmt.Println(y)

}
