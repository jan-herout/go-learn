package main

import "fmt"

// start,OMIT
func isTrue() bool { 
	fmt.Println("in a func")
	return true
}

func main() {
	x := true || isTrue()		// funkce se NEZAVOLÁ, short circuit
	fmt.Println(x)
}
// end,OMIT