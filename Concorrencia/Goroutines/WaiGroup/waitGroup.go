package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("Ola mundo")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Porgramando em Golang!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Dev Junior")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Dev Pleno")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() //0

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
