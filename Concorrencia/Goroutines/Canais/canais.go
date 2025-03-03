package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Ola Mundo!", canal)
	fmt.Println("Depois da funcao escrever comecar a ser executada")

	// for {
	// 	mensagem, aberto := <-canal
	// 	if !aberto {
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim da execucao")

}

func escrever(texto string, canal chan string) {
	// time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto // mandando valor para dentro do canal
		time.Sleep(time.Second)
	}

	close(canal)

}
