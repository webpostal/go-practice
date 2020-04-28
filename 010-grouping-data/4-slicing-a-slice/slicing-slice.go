package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}
	y := x[1:4]
	fmt.Printf("The length of the slice is %v\n", len(x))
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[2:])
	fmt.Println(x[1:3])
	fmt.Println(x[3:5])
	fmt.Printf("Printing a slice of x: %v\n", y)

	// Using range:

	for i, v := range x {
		fmt.Println(i, v)
	}

	// Printing slice without using range

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
