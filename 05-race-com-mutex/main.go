package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var result int
var m sync.Mutex

func main() {
	go runProcess("P1", 20)
	go runProcess("P2", 20)

	var s string
	fmt.Scanln(&s)
	fmt.Println("Final result: ", result) // 40
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)

		// TRAVA até UNLOCK para somente um processo
		// o processo que queira mexer em result deve aguardar o unlock
		m.Lock()
		result++
		fmt.Println(name, "->", i, "Partial result:", result)
		m.Unlock()

	}
}
