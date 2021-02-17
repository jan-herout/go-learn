package main

import (
	"fmt"
)

func main() {
	// START,OMIT
	work := make(chan string)
	go func() {
		<-work
		fmt.Println("DONE")
	}()
	close(work)
	work <- "data"
	// END,OMIT
}
