package main

//import "fmt"

func main() {
	// START,OMIT
	var m map[string]int       // nil
	m["abcd"] = 1              // PANIC
	// END,OMIT
}