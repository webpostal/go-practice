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
	fmt.Println(m["Tupac"])
	fmt.Println(m["Brad Meldhau"])

	delete(m, "Rianna")
	delete(m, "Snoop Dog")

	if v, ok := m["Tupac"]; ok {
		fmt.Println("value: ", v)
		delete(m, "Tupac")
	}

	fmt.Println(m)
}
