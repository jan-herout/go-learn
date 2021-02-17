package main

import (
	"fmt"
	"math/rand"
	"time"
)

// k funkcím se ještě dostaneme ....
func goOn() bool { return rand.Intn(100) < 80 }

func main() {
	// potřebujeme připravit zdroj entropie
	rand.Seed(time.Now().UnixNano())		

	
	// 80% pravděpodobnost, že budeme pokračovat
	var i int
	for goOn() {
		i++
		fmt.Println("tick")
	}
	fmt.Printf("iterations=%d\n",i)
}