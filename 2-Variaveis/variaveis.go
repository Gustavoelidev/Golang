package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	varivel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(varivel2)

	var (
		variavel4 string = "lala"
		variavel3 string = "lala"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel5", "Variavel6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
