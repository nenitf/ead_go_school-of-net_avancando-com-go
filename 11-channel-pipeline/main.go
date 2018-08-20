package main

/*
mind blowing

pipeline: um consumindo o outro

ORDEM:

generate -> divide
divide -> main

channel1 -> channel2
channel2 -> result

*/

import (
	"fmt"
)

// Pipeline
func main() {
	numbers := generate(2, 4, 6)

	// novo channel
	result := divide(numbers)

	// libera o channel result e aguarda até poder ser liberado novamente
	fmt.Println(<-result)

	// ja possui informação?
	fmt.Println(<-result)
	fmt.Println(<-result)
}

// func variatica, recebe infinitos int retorna chan int
func generate(numbers ...int) chan int {
	// cria channel inteiro
	channel := make(chan int)

	// go func que atribui valor ao channel
	go func() {
		for _, number := range numbers {
			// sempre que possivel, channel recebe um numero
			channel <- number
		}
	}()
	// retorna o channel sempre que possuir informação
	return channel
}

// func recebe chan int retorna chan int
func divide(input chan int) chan int {
	channel := make(chan int)

	// inicia a go func
	go func() {
		// cada vez que o input receber um numero - func generate
		// ele repassa a metade para o novo channel
		for number := range input {
			channel <- number / 2
		}
		// após dividir todos, bloqueia ESTE channel para evitar inserções
		close(channel)
	}()

	// retorna o channel sempre que ele estiver cheio
	return channel
}
