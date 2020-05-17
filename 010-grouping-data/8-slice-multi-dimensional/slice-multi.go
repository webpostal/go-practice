package main

import "fmt"

func main() {

	mj := []string{"Michael", "Jordan", "Bulls", "Shooting guard"}
	fmt.Println(mj)
	lb := []string{"Larry", "Bird", "Celtics", "Power forward"}
	fmt.Println(lb)
	sp := []string{"Scotty", "Pippen", "Bulls", "Small Forward"}

	lineup := [][]string{mj, lb, sp}
	fmt.Println(lineup)
}
