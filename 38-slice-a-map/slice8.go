package main

import "fmt"

func main() {
	// START,OMIT		
	var soi []int             // nil
	for i:= range soi {       // nil slice má kapacitu nula, a proto range *nezhavaruje*
		fmt.Println(i)
	}
	// END,OMIT
}
