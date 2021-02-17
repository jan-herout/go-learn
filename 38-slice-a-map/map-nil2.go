package main

import "fmt"

// START,OMIT
func checkMap(m map[string]int) {
	if m == nil {
		fmt.Println("map not initialized")
	}
	if len(m) == 0 {
		fmt.Println("map is empty (no members)")
	}
	fmt.Println("--------------------------------------")
}

func main() {
	var m map[string]int // nil

	checkMap(m)
	m = make(map[string]int)
	checkMap(m)
}
// END,OMIT
