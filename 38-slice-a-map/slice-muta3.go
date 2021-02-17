package main

import "fmt"

// START,OMIT		
func mutate(s []int) []int {	
	s[0] =  998
	s = append(s, 999)
	fmt.Printf("mutate end: %#v\n", s)
	return s
}

func main() {
	s := []int{1,2,3}
	fmt.Printf("main start: %#v\n", s)
	s = mutate(s)
	fmt.Printf("main end  : %#v\n", s)
}
// END,OMIT
