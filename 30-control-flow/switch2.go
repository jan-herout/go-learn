package main

import (
	"fmt"
	"time"
)

func main() {
// START,OMIT
	hour := time.Now().Hour() 
	switch {					// nebo: switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Máme dopoledne.")
	case hour < 17:
		fmt.Println("Příjemné odpoledne!")
	default:
		fmt.Println("Čas na večeči.")
	}
// END,OMIT
}