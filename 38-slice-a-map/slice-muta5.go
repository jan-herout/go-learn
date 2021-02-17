package main

import "fmt"

// START,OMIT		
func mutate(s []int) []int {	
	s[0] =  998
	s = append(s, 999)
	return s
}

func main() {
	s := []int{1,2,3}           
	sCopy := make([]int, len(s))
	copy(sCopy, s)
	
	fmt.Printf("    s[0] je na adrese %p\n", &s[0])     // oba slices "žijí" na jiných adresách
	fmt.Printf("sCopy[0] je na adrese %p\n", &sCopy[0]) // (myšleno "pole" jejich elementů)
	
	s = mutate(s)                                       // uprav původní slice
	fmt.Printf("copy is...: %#v\n", sCopy)              // kopie zůstala beze změny
}
// END,OMIT
