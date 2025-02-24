package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	TipoDeEndereco := enderecos.TipoDeEndereco("show Sao Francisco de Assis")
	fmt.Println(TipoDeEndereco)
}
