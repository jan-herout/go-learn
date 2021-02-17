package main

import "fmt"

// START,OMIT
// sum returns sum of all provided integer PLUS 10
func sum(integers ...int) int {                        // integers je variadický parametr
	var r int	
	fmt.Printf("in: %#v\n", integers)                  // interně je to vlastně slice
	integers = append(integers,10)                     // SYNTAKTICKY je korektní ho měnit
	fmt.Printf("changed: %#v\n", integers)
	for _, i := range integers {                       // a je možné ho procházet pomocí range
		r += i
	}
	return r
}

func main() { fmt.Println("součet je", sum(1,2,3,4,5))
	          fmt.Println("součet je",  sum()) }
// END,OMIT