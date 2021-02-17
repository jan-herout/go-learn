package main

import "fmt"

var aBool bool
var aString string
var aNumber int
var aPointer *int
var aSlice []string

func main() {
	// start,OMIT
	fmt.Printf("bool    = %#v\n", aBool)
	fmt.Printf("string  = %#v\n", aString)
	fmt.Printf("number  = %#v\n", aNumber)
	
	fmt.Printf("pointer = %#v\n", aPointer)
	fmt.Printf("slice   = %#v\n", aSlice)
	fmt.Printf("struct  = %#v\n", struct { name, surname string }{} )
	// end,OMIT
}