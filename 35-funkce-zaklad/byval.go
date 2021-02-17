package main

import "fmt"

// START,OMIT
func addToString(a,b string) string { // a i b jsou dočasné kopie vstupních hodnot
	a = a + " " + b 	// do a přiřadí konkatenaci
	return a			// a vrátí konkatenovanou hodnotu
}

func main() {
	s := "ahoj"
	fmt.Println(addToString(s,"lidi"))	// = ahoj lidi
	fmt.Println(s)						// = ahoj
}
// END,OMIT