package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
	height    float64
}

func main() {

	type secretAgent struct {
		person
		ltk bool
	}

	sa1 := secretAgent{
		person: person{
			firstname: "James",
			lastname:  "Bond",
			age:       32,
		},
		ltk: true,
	}

	p2 := person{

		firstname: "Nelson",
		lastname:  "Mandela",
		age:       95,
		height:    6.00,
	}

	fmt.Println(sa1)
	fmt.Println(p2)
	fmt.Println(sa1.firstname, sa1.lastname, sa1.age, sa1.ltk)
	fmt.Println(p2.firstname, p2.lastname, p2.age)
}
