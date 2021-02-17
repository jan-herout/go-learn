package main

import "fmt"

func main() {
	var ar [5]int
	// START,OMIT
	for i := range ar {                   // range vrací indexy pole, počínaje nejmenším
		ar[i] = i                         // prvku na indexu i přidělíme hodnotu i
		fmt.Println("i=",i, "val=",ar[i]) // prvky pole se indexují *od nuly!*
	}
	// END,OMIT
}