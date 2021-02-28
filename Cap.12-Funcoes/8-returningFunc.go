package main

import (
	"fmt"
)

func main() {

	x := retornaumafuncao()
	y := x(3)
	fmt.Println(y)

	//TAMBEM FUNCIONA DESSA FORMA
	fmt.Println(retornaumafuncao()(4))
}

func retornaumafuncao() func(int) int {
	return func(i int) int {
		return i * 11
	}
}
