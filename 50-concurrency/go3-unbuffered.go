package main

import (
	"fmt"
	"time"
)

func main() {
	// START,OMIT
	done := make(chan struct{}) // signál dokončení
	until := make(chan int)     // do jakého čísla chceme mít výpis?

	// ticker simuluje nějakou časově náročnou úlohu
	ticker := func(id string) {
		var max int
		// do max načti hodnotu z kanálu until - toto je blokující operace
		// ****DEADLOCK*** max = <-until
		for i := 0; i < max; i++ {
			fmt.Printf("ticker=%s, iter=%d\n", id, i)
			time.Sleep(time.Millisecond)
		}
		done <- struct{}{}
	}

	go ticker("alfa")    // běž na pozadí
	until <- 10          // BLOCKED - na kanálu není "nikdo", kdo by ho chtěl číst!
	<-done               // 
	fmt.Println("DONE!") // ********** DEADLOCK ****************
	// END,OMIT
}
