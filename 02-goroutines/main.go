package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// utilizar "go" na frente cria processo a parte
	// concorrencia é simultanea
	go runProcess("P1", 20)
	go runProcess("P2", 20)

	// cuida as alterações pelo runprocess
	// os processos de goroutines executam separadamente ao programa, logo eles seriam inependentes. é necessário ter uma "pausa" para ver os processos
	var s string
	fmt.Scanln(&s)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)

		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
}
