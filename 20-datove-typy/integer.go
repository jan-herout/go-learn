package main

import "fmt"

func main() {
	// nové proměnné
	var x, z, undef int // lze deklarovat víc proměnných najednou
	y := 0x00           // deklaruje proměnnou x, datový typ odvodí podle R-val, a inicializuje ji
	y, z = 2, 1         // atomická operace - přiřazení dvou hodnot najednou
	y, z = z, y         // atomická operace, prohození hodnot

	// QUIZ: co to vypíše?
	fmt.Printf("x=%d, y=%d, z=%d, undef=%d", x, y, z, undef)
}
