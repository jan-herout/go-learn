package main

import "fmt"

func main() {
	// START,OMIT
	//  DEKLARACE   | <-- INICIALIZACE -->
	var s1 []string = []string{"a", "b"}
	s2 := []string{"a", "b"}
	fmt.Printf("%#v, %#v\n", s1, s2)
	// END,OMIT
}
