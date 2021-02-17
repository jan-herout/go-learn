package main

import (
	"fmt"
	"strings"
)

func splitAndSend(s string) <-chan string {
	c := make(chan string)
	go func() {
		for _, w := range strings.Split(s, " ") {
			c <- w
		}
		close(c)
	}()
	return c
}

func main() {
	// START,OMIT
	for w := range splitAndSend("příšerně žluťoučký kůň úpěl ďábelské ódy") {
		fmt.Println(w)
	}
	// END,OMIT
}
