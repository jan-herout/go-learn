package main

//import "fmt"

func main() {
	var whatever interface{}
	whatever = "nějaký string"
	// START,OMIT
	whatever = whatever + " s pokračováním"
	// END,OMIT
	_ = whatever
}