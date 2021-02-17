package main

import (
	"fmt"
	"time"
)

func main() {
	// START,OMIT
	// ticker simuluje nějakou časově náročnou úlohu
	ticker := func(id string) {
		for i := 0; i < 100; i++ {
			fmt.Printf("ticker=%s, iter=%d\n", id, i)
			time.Sleep(time.Millisecond)
		}
	}
	// ticker je spuštěn na pozadí, ale *nečekáme* na jeho dokončení!
	go ticker("alfa")
	go ticker("omega")
	// tady je vidět, co se stane, pokud nesynchronizuji vstup a výstup
	time.Sleep(30 * time.Millisecond)
	fmt.Println("Už nečekáme!")
	// END,OMIT
}
