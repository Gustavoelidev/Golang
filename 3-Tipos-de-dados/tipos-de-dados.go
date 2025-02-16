package main

import (
	"errors"
	"fmt"
)

func main() {
	// int8, int16, int32, int64 /// tamanhos de bytes
	var numero int = 1000000000000000
	fmt.Println(numero)

	// uint //unsygned int
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	// Numeros reais

	var numeroReal1 float32 = 123.53
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123000000000000000000.53
	fmt.Println(numeroReal2)

	numeroreal3 := 12345.67
	fmt.Println(numeroreal3)

	// Fim numeros reais

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B' // tabelas ask
	fmt.Println(char)

	// FIM Strings

	texto := 5
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro) // nill = 0

}
