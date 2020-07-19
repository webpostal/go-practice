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

	/*fmt.Println(p1.firstname)
	fmt.Println(p1.lastname)
	for i, v := range p1.favIcecream {
		fmt.Println(i, v)
	}*/

	m := map[string]person{
		p1.lastname: p1,
		p2.lastname: p2,
	}

	/*fmt.Println(m)

	//Range over the Map

	for k, v := range m {
		fmt.Println(k, v)

	} */

	for _, v := range m {
		// fmt.Println(k, v)
		//fmt.Println(k)
		fmt.Println(v.firstname)
		fmt.Println(v.lastname)
		for i, val := range v.favIcecream {
			fmt.Println(i, val)
		}
		fmt.Println("---------")

	}
}
