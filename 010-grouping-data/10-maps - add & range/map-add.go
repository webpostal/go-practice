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
	fmt.Println(v, ok)

	// Add Key/Values:

	m["Leon"] = 43
	fmt.Println(m)
	fmt.Println(m["Leon"])

	// Range over keys:values

	for k, v := range m {
		fmt.Println(k, v)
	}

	// Same as range for slices Key = index, & the value for each key ->

	sli := []int{100, 200, 300, 400, 500, 600}

	for k, v := range sli {
		println(k, v)
	}

}
