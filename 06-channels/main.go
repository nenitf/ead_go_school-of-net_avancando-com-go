package main

import (
	"fmt"
)

func main() {

	// make cria slices, map e channels
	// channel do tipo string
	msg := make(chan string)

	// função anonima
	go func() {
		// mandar para o channel string "Hello World!"
		// channels recebem com "<-"
		msg <- "Hello World"
	}()

	// "transfere" o valor do channel para result com "<-"
	// channel é esvaziado e pode receber outro valor
	result := <-msg
	fmt.Println(result)

}
