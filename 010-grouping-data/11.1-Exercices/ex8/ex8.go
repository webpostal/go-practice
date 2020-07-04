package main

import "fmt"

func main() {

	m := map[string][]string{
		"Bunny Bugs":   []string{"Carrots", "Jokes", "Running"},
		"Coyote Wiley": []string{"Canyons", "Rockets", "Road Runners"},
		"Bob Sponge":   []string{"Krabby Patty", "Jelly Fish", "Karate"},
	}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for k2, v2 := range v {
			fmt.Println("\t", k2, v2)
		}
	}

}
