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
		max = <-until
		for i := 0; i < max; i++ {
			fmt.Printf("ticker=%s, iter=%d\n", id, i)
			time.Sleep(time.Millisecond)
		}
		done <- struct{}{}
	}

	go ticker("alfa") // běž na pozadí
	until <- 10       // maximálně 10 položek
	<-done            // počkej na dokončení
	fmt.Println("DONE!")
	// END,OMIT
}
