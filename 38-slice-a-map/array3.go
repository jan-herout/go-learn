package main

import "fmt"

func main() {
	var ar [5]int
	// START,OMIT
	fmt.Println(ar[0])  // první prvek má index 0
	fmt.Println(ar[4])  // tohle je poslední prvek pole
	fmt.Println(ar[5])  // tenhle prvek už neexistuje - COMPILE ERROR
	// END,OMIT
}