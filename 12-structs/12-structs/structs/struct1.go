package main

import "fmt"

func main() {
	type person struct {
		firstname string
		lastname  string
		age       int
		height    float64
	}
	p1 := person{
		firstname: "Raul",
		lastname:  "Castro",
		age:       78,
		height:    5.75,
	}

	p2 := person{
		firstname: "Nelson",
		lastname:  "Mandela",
		age:       95,
		height:    6.00,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.firstname, p1.lastname, p1.age)
	fmt.Println(p2.firstname, p2.lastname, p2.age)
}
