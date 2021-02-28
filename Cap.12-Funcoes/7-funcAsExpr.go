package main

import (
	"fmt"
)

func main() {
	x := 13

	y := func(x int) {
		fmt.Println(x, "vezes 223 eh:")
		fmt.Println(x * 223)
	}

	y(x)

}
