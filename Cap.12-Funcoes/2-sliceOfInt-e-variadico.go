package main

import (
	"fmt"
)

func main() {

	//sliceOfInt !=  int
	sliceOfInt := []int{10, 11, 12, 15}

			  //Qndo eh slice precisa passar com ... (variadico)
	totalSoma := soma(sliceOfInt...)
	fmt.Println(totalSoma)
}

func soma(y ...int) int {
	soma := 0
	for _, v := range y {
		soma += v
	}
	return soma
}
