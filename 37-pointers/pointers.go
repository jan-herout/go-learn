package main

import "fmt"
//START,OMIT
func byVal(a int)  {  a =  a + 1 }            // předej int "by value"
func byRef(a *int) { *a = *a + 1 }            // předej pointer na int "by value"
	
func main() {
	a := 0
	byVal(a)
	fmt.Println("var after by value", a)			// co to vrátí?
	byRef(&a)
	fmt.Println("var after by reference", a)		// co to vrátí?
}
//END,OMIT