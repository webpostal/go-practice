package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)

	}

	fmt.Printf("The slice 'x' is of type: %T\n\n", x)
	fmt.Println("Length is equals to:", len(x))
	fmt.Println("Capacity is equals to:", cap(x))

}
