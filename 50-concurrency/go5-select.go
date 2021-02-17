package main

import (
    "fmt"
    "time"
)

func generateData(c1, c2 chan string) {
	go func() {
        time.Sleep(2 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(1 * time.Second)
        c2 <- "two"
    }()
}

func main() {
	// START,OMIT
	// mějme dva odlišné procesy, ze kterých potřebujeme sebrat data,
	// a zpracovat je ASAP. Nevíme ale, který z procesů skončí dřív.
    c1 := make(chan string)
    c2 := make(chan string)
	
	// spouštíme zpracování na pozadí. Rozjedou se dva procesy, které
	// generují nezávisle na sobě data.
	generateData(c1,c2)			
    
	// víme, že čekáme na dvě položky, ale nevíme, co přijde dřív
    for i := 0; i < 2; i++ {
        // a protože čtení z kanálu je *blokující*, musíme se zkusit podívat
		// do obou kanálů; tohle je blokující read, přičemž ten kanál, který
		// poskytne data jako první, vyhraje.
		select {
        case msg1 := <-c1:
            fmt.Println("received from c1: ", msg1)
        case msg2 := <-c2:
            fmt.Println("received from c2: ", msg2)
        }
    }
	// END,OMIT
}