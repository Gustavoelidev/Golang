package main

import "fmt"

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan uint, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	// Consumir os resultados para evitar bloqueios
	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(tarefas <-chan int, resultados chan<- uint) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) uint {
	if posicao <= 1 {
		return uint(posicao)
	}
	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
