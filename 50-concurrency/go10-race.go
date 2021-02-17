package main

import (
	"fmt"
	"time"
)

func main() {
	// START,OMIT
	ints := make([]int, 100) // zero initialized slice: {0, 0, .... 0}
	sumChan := make(chan int)

	sumFunc := func() {
		sum := 0
		for _, i := range ints {
			sum += i
		}
		sumChan <- sum
	}

	setFunc := func() {
		for i := 99; i >= 0; i-- {
			ints[i] = 1
			time.Sleep(time.Microsecond * 4)
		}
	}

	go sumFunc()
	go setFunc()
	sum := <-sumChan
	fmt.Println("součet by měl být 0, ale je ", sum)
	// END,OMIT
}
