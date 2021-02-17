package main

import "fmt"

var balance = 100

func main() {
	fmt.Println("puvodni castka: ", balance)
	if true {
		balance := balance + 100 					// shadowing
		fmt.Println("pridame 100, mame", balance)
	}
	fmt.Println("nova castka: ", balance) 			// OPPS!
}
