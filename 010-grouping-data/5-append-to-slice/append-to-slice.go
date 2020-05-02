package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 111)
	fmt.Println(x)

	y := []int{123, 456, 789, 987}
	x = append(x, y...)
	fmt.Println(x)

}
