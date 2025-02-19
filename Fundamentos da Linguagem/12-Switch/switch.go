package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Quarta"
	case 4:
		return "Quinta"
	case 5:
		return "Sexta"
	case 6:
		return "Sabado"
	case 7:
		return "Domingo"
	default:
		return "Numero invalido"
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	default:
		return "Numero invalido"
	}
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(1)
	fmt.Println(dia)
	fmt.Println("###### Novo teste #########")
	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
