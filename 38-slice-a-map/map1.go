package main

import "fmt"

func main() {
	// START,OMIT
	m := make(map[string]int)
	m["abcd"] = 1
	m["xyz"] = 2
	fmt.Println("%#v\n",m)
	// END,OMIT
}