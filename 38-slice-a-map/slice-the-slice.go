package main

import "fmt"

func main() {
// START,OMIT		
	s1 := []int{1,2,3,4,5}
	s2 := s1[1:3]				// s2 = []int{2,3}      // INCLUSIVE:EXCLUSIVE

	// s1 a s2 mají různou délku
	fmt.Printf("s1=%#v\ns2=%#v\n", s1, s2)
	
	// ovšem pozor, s2 je "řez" z s1!
	s2[0] = 100
	fmt.Printf("&s1[1]=%p\n&s2[0]=%p\n", &s1[1], &s2[0])
	fmt.Printf("s1=%#v\ns2=%#v\n", s1, s2)
	
	// tohle jde taky udělat... všechno kromě posledního prvku
	s2 = s1[ : len(s1) - 1]
	fmt.Printf("s1=%#v\ns2=%#v\n", s1, s2)
// END,OMIT
}
