package main

import "fmt"

func main() {
	// START,OMIT
	var s []string // slice typu string
	fmt.Printf("%#v: len=%d, cap=%d\n", s, len(s), cap(s))
	if s == nil { fmt.Println("zero value pro slice je nil!") }
	// END,OMIT
}
