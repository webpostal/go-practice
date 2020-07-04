package main

import "fmt"

func main() {

	m := map[string][]string{
		"Bunny Bugs":   []string{"Carrots", "Jokes", "Running"},
		"Coyote Wiley": []string{"Canyons", "Rockets", "Road Runners"},
		"Bob Sponge":   []string{"Krabby Patty", "Jelly Fish", "Karate"},
	}

	m["Piggy Miss"] = []string{"Frogs", "Hairstyles", "Sweets"}

	fmt.Printf("%v\n\n", m)

	// delete(m, "Bunny Bugs")

	for k, v := range m {
		fmt.Printf("This is the record corresponding to: %v\n", k)
		for _, v2 := range v {
			fmt.Println("\t", v2)
		}
	}
	anotherMapFunc()
}

func anotherMapFunc() {
	y := map[int]string{1: "Money", 2: "Health", 3: "Education"}
	fmt.Println(y)

	y[4] = "Virtue"
	for t, val := range y {
		fmt.Println(t, val)
	}

}
