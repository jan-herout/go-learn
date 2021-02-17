package main

import "fmt"

// START,OMIT		
func mutate(s []int) []int {	
	s[0] =  998
	s = append(s, 999)
	return s
}

func main() {
	s := []int{1,2,3}                      // inicializace hodnotami 1,2,3
	sCopy := s                             // vytvoř si kopii bokem
	s = mutate(s)                          // uprav původní slice
	fmt.Printf("copy is...: %#v\n", sCopy) // ???HUH ???
}
// END,OMIT
