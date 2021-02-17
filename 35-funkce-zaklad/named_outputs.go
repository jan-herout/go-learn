package main

import "fmt"

// START,OMIT
func greeter() (name, greetings string) {
	// name a greetings jsou deklarované v signatuře funkce
	// není tedy možné je deklarovat znovu	
	name = "Prostetnik Vogon Jelc"
	greetings = " Ó fretná chrochtobuznosti..."
	
	// Kompilátor ví, že vracíme dva parametry, a zná dopředu jejich jméno.
	// Není tedy syntakticky nutné je uvádět (ale POZOR, znepřehledňuje to kód!
	return		// totéž jako return name, greetings
}

func main() {
	n,g := greeter()
	fmt.Println(n, ":", g)
}
// END,OMIT