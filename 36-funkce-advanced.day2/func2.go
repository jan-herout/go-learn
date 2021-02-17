package main

import "fmt"



// START,OMIT
func main() {
	for i := 0; i<5; i++ {
		fmt.Println("i=",i)
		x := i
		defer func() { 
			fmt.Println("defered i=",i,"x=",x) 
		}()
	}	
// END,OMIT
}