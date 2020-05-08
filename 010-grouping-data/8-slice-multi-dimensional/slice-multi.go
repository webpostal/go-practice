package main

import "fmt"

func main() {

	mj := []string{"Michael", "Jordan", "Bulls", "Shooting guard"}
	fmt.Println(mj)
	lb := []string{"Larry", "Bird", "Celtics", "Power forward"}
	fmt.Println(lb)

	pl := [][]string{mj, lb}
	fmt.Println(pl)
}
