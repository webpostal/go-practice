package main

import "fmt"

func main() {

	x := make([]int, 10, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x[1] = 11
	x[9] = 1
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 45)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
