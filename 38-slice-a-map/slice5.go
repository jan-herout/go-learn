package main

import "fmt"

func main() {
	// START,OMIT	
	// slice s předenm danou délkou
	s1 := make([]string, 5)
	s1[0] = "abcd"
	s1[4] = "xyz"
	s1[5] = "#$#@%!!"			// oops!
	fmt.Println(s1)
	// END,OMIT
	_ = s1
}
