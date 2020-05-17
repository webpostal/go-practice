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

	m2 := map[int]string{
		1: "Monica",
		2: "Laura",
		3: "Martha",
	}

	fmt.Println(m2[3])

	if val, ok := m2[6]; ok {
		fmt.Println(val)

	} else {
		fmt.Printf("The value is %v \n", ok)
	}

}
