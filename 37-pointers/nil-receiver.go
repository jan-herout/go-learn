package main

import "fmt"

// START,OMIT
type myInt int
func (i *myInt) Greet() { fmt.Println("greetings") }
func main() {
	var pi *myInt            // nil
	fmt.Println(pi)
	pi.Greet()
	// END,OMIT
}