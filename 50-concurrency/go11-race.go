package main

import (
	"fmt"
	"sync"
)

func main() {
	// START,OMIT
	m := make(map[int]string)
	var wg sync.WaitGroup
	wg.Add(3)
	// PRODUCER 1
	go func() {
		m[0] = "ANO"
		wg.Done()
	}()
	// PRODUCER 2
	x := 0
	go func() {
		x = x + 1
		m[0] = "NE"
		wg.Done()
	}()
	// CONSUMER
	go func() {
		i := 0
		for i < 5 {
			fmt.Println(i, m[0])
			i++
		}
		wg.Done()
	}()
	wg.Wait()
	// END,OMIT
	_ = x
}
