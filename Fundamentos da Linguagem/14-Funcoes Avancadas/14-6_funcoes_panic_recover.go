package main

import "fmt"

func recuperar() {
	if r := recover(); r != nil {
		fmt.Println("execucao recuperada com sucesso!")
	}
}

func aluno(n1, n2 float64) bool {
	defer recuperar()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A MEDIA E EXATAMENTE ESTAA!!!")

}

func main() {
	fmt.Println(aluno(6, 6))
	fmt.Println("Pos execucao")
}
