package main

import "fmt"

func main() {
	var pi *int  
	// START,OMIT
	if pi == nil { 
		fmt.Println("pi je", pi) 
	}
	// END,OMIT
}