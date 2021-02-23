package main

import (
	"fmt"
)

func main() {

	basica()

	recebeArg("tarde")

	recebeRetorno := soma(10, 10)
	fmt.Println(recebeRetorno)

	totalSoma, qntdElem := somaeqntd(10, 10, 12)
	fmt.Println(totalSoma, qntdElem)
}

func basica() {
	fmt.Println("Bom dia")
}
func recebeArg(s string) {
	fmt.Println("Boa " + s)
}
func soma(d, e int) int {
	return d + e
}
func somaeqntd(y ...int) (int, int) {
	soma := 0
	for _, v := range y {
		soma += v
	}
	return soma, len(y)
}
