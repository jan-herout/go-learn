package main

import "fmt"

func main() {
	ar := [5]int{0,1,2,3,4}
	// START,OMIT
	for i, v := range ar {
		fmt.Printf("i=%d,v=%d\n",i,v)
		ar[i] = 10     // změní pole
		v     = 1000   // nemá vliv na "outer scope", mění kopii hodnoty v poli
	}
	fmt.Println(ar)
	// END,OMIT
}