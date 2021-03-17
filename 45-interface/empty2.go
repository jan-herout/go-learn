package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// START,OMIT
	check := func(s string, i interface{}) { if i == nil { fmt.Println(s,"NIL!" )} }	
	var r io.Reader
	var h *os.File
	check("empty interface", r)
	check("empty pointer", r)
	r = h
	check("interface s empty pointerem", r)	
	// END,OMIT
}