package main

import "fmt"

func main() {
	// func() {
	// 	fmt.Println("Ola universo!")
	// }()
	// se eu quiser rodar uma funcao sem name, somente colocar apos as chaves o ()
	// func(texto string) {
	// 	fmt.Println(texto)
	// }("Passando Parametro")

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parametro")

	fmt.Println(retorno)
}
