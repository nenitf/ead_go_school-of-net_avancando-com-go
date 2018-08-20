package main

/*

race conditions:
ambos processos utilizam a mesma variavel global
um processo interfere no outro

*/

import (
	"fmt"
	"math/rand"
	"time"
)

var result int

func main() {
	go runProcess("P1", 20)
	go runProcess("P2", 20)

	var s string
	fmt.Scanln(&s)
	fmt.Println("Final result: ", result) // 20, deveria ser 40!

	/*
		Para checar se esta ocorrendo uma race pode ser enviado para o terminal o comando:
		go run -race main.go
	*/
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		z := result
		z++
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		result = z
		fmt.Println(name, "->", i, "Partial result:", result)
	}
}
