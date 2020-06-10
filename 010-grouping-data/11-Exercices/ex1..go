package main

import "fmt"

func main() {
	m := [5]int{1, 5, 7, 9, 11}
	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)

	}

	fmt.Printf("The Array 'm' is of type: %T\n\n", m)

	// Default:Zero value array
	t := [5]int{}
	fmt.Println(t)

}
