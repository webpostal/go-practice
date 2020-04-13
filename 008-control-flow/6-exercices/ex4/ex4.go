package main

import "fmt"

func main() {

	bd := 1974
	yr := 2020

	for {
		if bd > yr {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
