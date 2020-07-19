package main

import "fmt"

func main() {

	myStruct := struct {
		name   string
		colors []string
		scores map[string]int
	}{
		name:   "Marco Polo",
		colors: []string{"red", "green", "blue"},
		scores: map[string]int{
			"1998": 154,
			"2005": 205,
		},
	}
	fmt.Println(myStruct)
	fmt.Println(myStruct.name)
	fmt.Println(myStruct.colors)
	fmt.Println(myStruct.scores)
	for k, val := range myStruct.colors {
		fmt.Println(k, val)
	}
	for i, v := range myStruct.scores {
		fmt.Println(i, v)

	}

}
