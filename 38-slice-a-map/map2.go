package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["abcd"] = 1
	m["xyz"] = 2

	// START,OMIT
	for k,v := range m {
		fmt.Printf("ORIG:  key=%s, val=%d\n",k,v)
		m[k]++		// upraví obsah mapy
		fmt.Printf("NEW :  key=%s, val=%d\n",k,m[k])
		fmt.Println("--------------------------------------")
	}
	// END,OMIT
}