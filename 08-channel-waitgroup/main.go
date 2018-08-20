package main

import (
	"fmt"
	"sync"
)

func main() {

	// cria channel que recebe somente int
	channel := make(chan int)

	// cria waitGroup
	var waitGroup sync.WaitGroup

	// especifica que existem dois waitGroups existentes
	waitGroup.Add(2)

	// func paralela anonima que acrescenta ao channel
	go func() {
		for i := 0; i < 10; i++ {
			// channel recebe SE estiver vazio para isso
			channel <- i
		}
		// após terminar de inserir os 10 valores, maraca UM done
		waitGroup.Done()
	}()

	// mesma func como a de cima
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		waitGroup.Done()
	}()

	// func paralela aninoma que fica "escutando" os groupWait.Done
	// após TODOS dones - 2 - o channel é encerrado
	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	// main
	// toda vez que numbermber recebe channel o esvazia
	for number := range channel {
		fmt.Println(number)
	}

}
