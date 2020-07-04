package main

import "fmt"

func main() {
	sx1 := []string{"Michael", "Jordan", "The last dance"}
	sx2 := []string{"Quentin", "Tarantino", "Pulp Fiction"}
	fmt.Println(sx1)
	fmt.Println(sx2)

	ss1 := [][]string{sx1, sx2}
	fmt.Println(ss1)

	for i, sx := range ss1 {
		fmt.Println("record: ", i)
		for j, val := range sx {
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
