package main

import "fmt"

func fact(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}	
	return n * fact(n-1)   // rekurze ... což možná není nejlepší řešení
}

func main() {
	fmt.Println(fact(10))
}