package main

import "fmt"

func main() {
	// START,OMIT
	var pi *int            // nil
	fmt.Println(*pi)       // dereferencuj nil => PANIC!
	// END,OMIT
}