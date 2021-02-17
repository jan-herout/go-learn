package main

import "fmt"

func main() {
	// START,OMIT	
	// slice s předenm danou délkou
	s1 := make([]string, 0)
	fmt.Printf("%#v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	
	// slice s předem danou délkou a kapacitou
	wantLen, wantCap := 2,10
	s2 := make([]int, wantLen, wantCap)
	fmt.Printf("%#v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
	// END,OMIT
}
