package main

/*
uma função paralela espera pela outra
pelo fato dos channels possuirem a regra de só receber um valor por vez


exemplo: corrida e bastão
*/

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int)

	// função anonima paralela
	// altera o valor do channel
	go func() {
		for i := 0; i < 10; i++ {
			// o loop aguarda até que channel possa receber i
			// ou seja, intercala com a função de imprimir e remover
			channel <- i
		}
	}()

	// func imprime channel
	// func retira "<-" o valor de channel
	go func() {
		for {
			fmt.Println(<-channel)
		}
	}()

	// pausa a função
	time.Sleep(time.Second)

}
