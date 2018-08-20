package main

/*
Pensar em semaforo com sua sinalização sendo flags

É uma alternativa ao waitGroup
*/

import (
	"fmt"
)

func main() {
	// dois channels
	channel := make(chan int)
	ok := make(chan bool) // flag

	// sempre que puder channel recebe um valor
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		// flag
		ok <- true
	}()

	// func duplicata
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		// flag
		ok <- true
	}()

	go func() {
		// vai apagar ok duas vezes antes de fechar o channel
		// flags antes de encerrar channel
		<-ok
		<-ok
		close(channel)
	}()

	// vai lendo enquanto existir valor em channel
	// retira o valor sempre que é lido
	for number := range channel {
		fmt.Println(number)
	}
}
