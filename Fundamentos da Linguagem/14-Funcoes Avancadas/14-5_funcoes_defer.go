package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a funcao 1")
}

func funcao2() {
	fmt.Println("Executando a funcao 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media Calculada - Resultado sendo gerado!")
	fmt.Println("Entrando na funcao para verificar se o aluno esta aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	// defer funcao1()
	// funcao2()
	fmt.Println(alunoEstaAprovado(7, 8))
}

// Defer = Adiar, ele ira adiar a funcao ate o ultimo momento possivel.
