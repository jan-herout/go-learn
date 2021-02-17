package main

import "fmt"

func main() {
	var x float64 = 3.0
	y := 2             	
	// .\float-bug.go:8:16: invalid operation: x / y (mismatched types float64 and int)
	fmt.Println(x / y) 
	// tohle ale projde
	fmt.Println(x/float64(y))
}
