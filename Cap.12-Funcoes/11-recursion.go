package main

import (
	"fmt"
)

func main() {

	fmt.Println(fatorialRecursivo(4))
	fmt.Println(fatorialLoop(4))

}

func fatorialRecursivo(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorialRecursivo(x-1)
}

func fatorialLoop(x int) int {
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}
