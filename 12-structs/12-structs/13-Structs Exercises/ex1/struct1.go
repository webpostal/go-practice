package main

import "fmt"

type person struct {
	firstname   string
	lastname    string
	favIcecream []string
}

func main() {

	p1 := person{
		firstname:   "Michael",
		lastname:    "Bolton",
		favIcecream: []string{"Bubble Gum", "Chocolate Chip", "Vanilla Ice"},
	}

	p2 := person{
		firstname:   "Laura",
		lastname:    "Mirkovic",
		favIcecream: []string{"Soda Pop", "Mango Delight", "Dulce de Leche"},
	}

	fmt.Println(p1.firstname)
	fmt.Println(p1.lastname)
	for i, v := range p1.favIcecream {
		fmt.Println(i, v)
	}

	fmt.Println(p2.firstname)
	fmt.Println(p2.lastname)
	for i, v := range p2.favIcecream {
		fmt.Println(i, v)
	}

}
