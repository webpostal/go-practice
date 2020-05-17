package main

import "fmt"

func main() {
	m := map[string]int{
		"Tupac":  32,
		"Sting":  55,
		"Bono":   53,
		"Rianna": 28,
	}
	fmt.Println(m)
	fmt.Println(m["Bono"])
	fmt.Println(m["Morrissey"])

	//Check if value exits

	v, ok := m["Morrissey"]
	//fmt.Println(v)
	fmt.Println(v, ok)

	m["Leon"] = 43
	fmt.Println(m["Leon"])

}
