package main

import (
	"fmt"
)

func main() {

	for i := 33; i <= 122; i++ {

		fmt.Printf("Unicode: %#U\tdecimal: %d\tHex: %#x\t String: %s\n", i, i, i, string(i))

	}

	/*
		Ejemplo usando un slice of byte []byte:

		s := "La verdad os hará libres"
		d := []byte(s)

		for i := 0; i < len(d); i++ {

			//fmt.Printf("%#U\t %s\n", d[i], string(d[i]))
			fmt.Printf("%s", string(d[i]))

		}
	*/

}
