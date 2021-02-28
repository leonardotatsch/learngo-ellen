package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salario          float64
}

type arquiteto struct {
	pessoa
	tipodeconstrucao string
	softwarequeedita string
}

func (j dentista) oibomdia() {
	fmt.Println("Meu nome eh:", j.nome, ", sou dentista e ja arranquei", j.dentesarrancados, "dentes")
}

func (q arquiteto) oibomdia() {
	fmt.Println("Meu nome eh:", q.nome, "sou arquiteto e trabalho com ", q.tipodeconstrucao)
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
	switch g.(type) {
	case dentista:
		fmt.Println("Eu ganho:", g.(dentista).salario)
	case arquiteto:
		fmt.Println("Eu edito com:", g.(arquiteto).softwarequeedita)
	}
}

func main() {
	srdente := dentista{
		pessoa: pessoa{
			nome:      "Arnaldo",
			sobrenome: "Birilimba",
			idade:     37,
		},
		dentesarrancados: 3448,
		salario:          23012.56,
	}
	srescadas := arquiteto{
		pessoa: pessoa{
			nome:      "Tonho",
			sobrenome: "Saldanha",
			idade:     68,
		},
		tipodeconstrucao: "escadaria",
		softwarequeedita: "nanquin",
	}
	serhumano(srdente)
	fmt.Println("")
	serhumano(srescadas)
}
