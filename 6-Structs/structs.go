package main

import "fmt"

type endereco struct {
	logadouro string
	numero    uint8
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Gustavo"
	u.idade = 22
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 0}

	usuario2 := usuario{"Gustavo", 22, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 22}
	fmt.Println(usuario3)
}
