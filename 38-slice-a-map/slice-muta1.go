package main

import "fmt"

// START,OMIT		
func mutate(s []int) {	
	s[0] =  998
	s = append(s, 999)
	fmt.Printf("mutate end: %#v\n", s)
}

func main() {
	s := []int{1,2,3}
	fmt.Printf("main start: %#v\n", s)
	mutate(s)
	fmt.Printf("main end  : %#v\n", s)
}
// END,OMIT
