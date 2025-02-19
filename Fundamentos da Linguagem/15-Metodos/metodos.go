package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando dados do Usuario %s no banco de dados\n", u.nome)
}

func (u usuario) idade18() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"usuario 1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()
	usuario2 := usuario{"Gustavo", 30}
	usuario2.salvar()
	idade18 := usuario2.idade18()
	fmt.Println(idade18)
	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
