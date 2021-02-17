package main

import (
	"fmt"
)

func main() {
	// SHOW OMIT
	var c3, c5 int
	for i:=1;i<16;i++ {
		out := ""
		c3, c5 = c3+1, c5+1		
		if c3 == 3 {
			c3, out = 0, "fizz"
		}
		if c5 == 5 { 
			c5, out = 0, out + "buzz"
		}
		if out == "" {
			fmt.Println(i)
			continue
		}
		fmt.Println(out)
	}
	// ENDSHOW OMIT
}
