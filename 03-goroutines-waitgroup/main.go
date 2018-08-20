package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

// func init executa antes do main
// utilizada para setup
func init() {
	// setar a quantidade de max cpus
	// o parametro setado é a qtd max da maquina
	// meu notebook possui 4 por exemplo
	runtime.GOMAXPROCS(runtime.NumCPU())

	/*
		concorrente: apenas 1 cpu entre processos
		paralelismo: mais de 1 cpu entre processos
	*/
}

func main() {
	//fmt.Println(runtime.NumCPU()) // qtd de cpus da maquina

	// "EXISTE UM WAIT GROUP AGUARDANO DUAS FINALIZAÇÕES"
	waitGroup.Add(2)
	go runProcess("P1", 20)
	go runProcess("P2", 20)

	// o script main só vai finalizar após todos "DONE" adicionados no  - 2
	waitGroup.Wait()
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, "->", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}

	// finaliza um wait group
	waitGroup.Done()
}
