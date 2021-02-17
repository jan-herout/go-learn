package main

import "fmt"

func main() {
	// START		
	s := make([]int, 0) // dopředu nemám žádnou indikaci o potřebné velikosti, proto len=cap=0
	fmt.Printf("po deklaraci: %#v, len=%d, cap=%d\n", s, len(s), cap(s))
	
	for i:=0; i<6; i++ {
		s = append(s, i)	// append PŘIDÁ nový prvek do slice
		fmt.Printf("append #%d   : len=%d, cap=%d, val=%#v\n", i, len(s), cap(s), s)
	}
	// END,OMIT
}
