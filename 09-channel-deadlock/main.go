package main

/*
SEMPRE atribuir valor a um channel em goroutine - concorrente ou paralela
*/

import (
	"fmt"
)

func main() {
	channel := make(chan int)
	// channel <- 10 // Deadlock!
	go func() {
		channel <- 10
	}()

	fmt.Println(<-channel)
}
