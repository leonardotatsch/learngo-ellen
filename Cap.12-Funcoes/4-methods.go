package main

import (
	"fmt"
)

type pessoa struct {
	nome  string
	idade int
}

// func (receiver) identifier(parameters) (returns) {code}
func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "diz bom dia!")
}

func main() {
	leo := pessoa{"Leonardo", 31}
	leo.oibomdia()
}

