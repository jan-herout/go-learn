package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func sleep(i int) {
	n := rand.Intn(5) + 1
	fmt.Printf("%d: will take just a few (%d) seconds....\n", i, n)
	time.Sleep(time.Duration(n) * time.Second)
}

var maxCnt int = 4

func main() {
	rand.Seed(time.Now().UnixNano())
	// proč pointer? protože dole je closure!
	var wg sync.WaitGroup
	work := make(chan int, maxCnt)

	writer := func(i int) {
		sleep(i)
		work <- i
		wg.Done()
	}

	closer := func() {
		wg.Wait()
		close(work)
	}

	readerDone := make(chan struct{})
	reader := func() {
		for i := range work {
			fmt.Println("received", i)
		}
		close(readerDone)
	}

	// START,OMIT
	go reader()                    // reader, konzument dat (na pozadí)
	for i := 0; i < maxCnt; i++ {  // několik "producentů" dat
		wg.Add(1)                  // DŮLEŽITÉ! wg.Add(1) *musí* proběhnout zde!
		go writer(i)               // run producer
	}
	go closer()   // odpal komponentu, která čeká na dokončení producentů
	<-readerDone  // a nakonec, počkej až skončí konzument ...
	// END,OMIT
}
