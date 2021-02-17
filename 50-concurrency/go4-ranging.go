package main

import (
	"fmt"
)

func main() {

	// START,OMIT
	// work je channel, na který budeme zasílat jednotlivé prvky
	work := make(chan string)

	// worker je funkce, která na pozadí spustí odběr a zpracování dat,
	// a vrátí kanál signalizující dokončení.
	// poté, co již není na vstupu nic ke zpracování, pošle signál ukončení.
	worker := func(c <-chan string) <-chan struct{} {
		done := make(chan struct{})
		go func() {
			for s := range c {  // dokud máš neco na vstupu
				fmt.Println(s)  // pracuj
			}                   // 
			close(done)         // a když tam už nic není, signalizuj konec
		}()
		return done
	}
	
	done := worker(work)	// spusť worker rutinu, a získej signalizační kanál
	work <- "řádek 1"		// rutina worker na pozadí spustila anonymní funkci
	work <- "řádek 2"       // která zpracovává data ....
	close(work)             // worker se musí NĚJAK dozvědět, že další data už nebudou
	<-done                  // worker uzavře signalizační kanál, a tím indikuje konec

	// END,OMIT
}
