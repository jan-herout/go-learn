package main

import (
    "fmt"
    "time"
)

func sendAfter(t time.Duration, c chan string, i string) {
	fmt.Println("sendAfter: sleeping... ")
	time.Sleep(t)
	fmt.Println("sendAfter: sending",i)
	c <- i
}

func main() {
	// START,OMIT
    work := make(chan string)
	// worker accepts work on string channel, and returns a done channel.
    worker := func(c <- chan string) <- chan struct{} {		
		done := make(chan struct{})
		go func() {                                          // na pozadí
			loop: for {                                      // máme nekonečnou smyčku
			select {                                         // 
				case work, more := <- c:                     // máme nějaká data?
					fmt.Printf("máme vstup [%s]\n", work)    //
					if !more { break loop }                  // a budou další data?
				case <- time.Tick(time.Second):              //--------------------
					fmt.Println("tick...")
				// default: fmt.Println("TOCK...")
			}                                                
			}
			fmt.Println("no more work!")
			close (done)
		}()
		return done
	}
	done := worker(work)
	sendAfter(2 * time.Second, work, "item 1")
	sendAfter(2 * time.Second, work, "item 2")
	close(work)
    <- done
	
	// END,OMIT
}