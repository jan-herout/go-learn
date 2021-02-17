package main

import "fmt"

// START,OMIT		
// mutate dostane na vstupu kopii struktury, která obsahuje:
//   - pointer na jednotlivé položky
//   - délku a kapacitu (int)
// předání proběhne by value (stack); pozor, součástí struktury je pointer!
func mutate(s []int) {
	// tahle změna OVLIVNÍ nadřazený scope, protože "dereferencujeme pointer"
	// a ovlivňujeme obsah dané adresy
	s[0] =  998
	// *tahle* změna ale mění délku a kapacitu *kopie* struktury slice
	//               a proto *NEOVLIVNÍ* nadřazený scope (není persistentní!)
	s = append(s, 999)
	fmt.Printf("mutate end: %#v\n", s)
}

func main() {
	s := []int{1,2,3}
	fmt.Printf("main start: %#v\n", s)
	mutate(s)
	fmt.Printf("main end  : %#v\n", s)
}
// END,OMIT
