package main

import (
	"fmt"
)

func main() {
	//START,OMIT
	s:="abcd"
	fmt.Println(s[0], s[1], s[2])
	s[1] = "X"	
	//END,OMIT
}
