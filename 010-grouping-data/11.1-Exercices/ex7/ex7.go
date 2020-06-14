package main

import "fmt"

func main() {
	xs1 := []string{"Michael", "Jordan", "The last dance"}
	xs2 := []string{"Quentin", "Tarantino", "Pulp Fiction"}
	fmt.Println(xs1)
	fmt.Println(xs2)

	ss1 := [][]string{xs1, xs2}
	fmt.Println(ss1)

	for i, xs := range ss1 {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}

	/*	ss2 := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		fmt.Println(ss2)*/

}
