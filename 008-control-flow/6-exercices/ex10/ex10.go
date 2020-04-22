package main

import "fmt"

func main() {

	fmt.Printf("true && true is %v\n", true && true)
	fmt.Printf("true && false is %v\n", true && false)
	fmt.Printf("false && true is %v\n", false && true)
	fmt.Printf("false && false is %v\n\n", false && false)

	fmt.Printf("true || true is %v\n", true || true)
	fmt.Printf("true || false is %v\n", true || false)
	fmt.Printf("false || true is %v\n", false || true)
	fmt.Printf("false || false is %v\n\n", false || false)

	fmt.Println(!false && !false)
}
