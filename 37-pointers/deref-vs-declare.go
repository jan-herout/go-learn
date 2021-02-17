package main

import "fmt"

func main() {
	// START,OMIT
	i := 0                             // deklaruje nový int
	var pi *int                        // deklaruje pointer
	pi = &i                            // reference - pi nyní odkazuje na hodnotu v proměnné i
	*pi = 1                            // měníme obsah hodnoty na adrese pi
	fmt.Println("i=",i,"*pi=",*pi)     // další příklad dereference
	// END,OMIT
}