package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Estrutura de Loops")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Gustavo", "Audi", "Arthur"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"Nome":      "Leonardo",
		"Sobrenome": "Silva",
	}
	// fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	type usuarioStruct struct {
		nome      string
		Sobrenome string
	}

	usuario2 := usuarioStruct{"Ze", "Guedes"}
	fmt.Println(usuario2)

	for {
		fmt.Println("Executando infinitamente....")
		time.Sleep(time.Second)
	}
}
