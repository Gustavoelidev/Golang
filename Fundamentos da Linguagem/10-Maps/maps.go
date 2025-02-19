package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{ // Limitacao do MAP o valor tem que ser coerentes.
		"Nome":      "Gustavo",
		"Sobrenome": "Eli",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"Primeiro nome": "Arthur",
			"Sobrenome":     "Eli",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Biguacu",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

}
