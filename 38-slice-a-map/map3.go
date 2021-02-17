package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["abcd"] = 1
	m["xyz"] = 2

	// START,OMIT
	k:= "abcd"
	if v, ok := m[k]; ok {
		fmt.Printf("klíč %s v mapě je, hodnota je %d, proměnná ok je %b", k, v, ok)
	}
	
	k= "tohle tam nenajdeš"
	if v, ok := m[k]; ok {
		// tohle se nikdy nespustí, protože ok == false
		fmt.Printf("klíč %s v mapě je, hodnota je %d, proměnná ok je %b", k, v, ok)
	} 
	// END,OMIT
}