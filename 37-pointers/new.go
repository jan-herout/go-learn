package main

import "fmt"

func main() {
	// START,OMIT
	p := new(string)                    // new: alokuje prostor pro nový string...
	fmt.Println("adresa je", p)         //      a vrátí jeho adresu
	
	*p = "text"                         // DEREFERENCE: změň hodnotu na adrese p
	fmt.Println("na adrese máme", *p)
	
	//-----------------------------------------------------------------------------------
	// stejné chování, ale čistě pomocí operátorů
	s := "text"                         // alokujeme prostor pro nový string
	ps := &s                            // REFERENCE: jeho adresu uložíme do proměnné ps
	fmt.Println("adresa:",ps,           
	            "hodnota na ní:", *ps)  // DEREFERENCE: co je na adrese ps?
	// END,OMIT
}