package main

import (
	"fmt"
	"time"
)

func main() {
	// START, OMIT	
	switch time.Now().Weekday() {
	case time.Saturday:
		fmt.Println("Dnes je sobota.")
	case time.Sunday:
		fmt.Println("Dnes je neděle.")
	default:
		fmt.Println("Hurá do práce!")
	}	
	// END, OMIT
}