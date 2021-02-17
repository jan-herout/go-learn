package main

import "fmt"

func main() {
	// START,OMIT
	m := make(map[string]int)
	fmt.Println(m["neexistující prvek"])
	// END,OMIT
}