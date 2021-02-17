package main

import "fmt"

// START,OMIT
type Person struct {
	Name     string
	YearBorn int
}

// NewPerson je vlastně KONSTRUKTOR "třídy" Person
func NewPerson(name string, yearBorn int) Person {
	p := Person{ // Tento zápis zakládá instanci struktury Person
		Name:     name,     // a nastavuje její jednotlivá pole.
		YearBorn: yearBorn, // Zápis je v podobě `pole:hodnota`.
	}
	return p // Konstruktor vrací *kopii!* instance p
}

func main() {
	// použití konstruktoru...
	p := NewPerson("John Doe", 1977)
	// příklad toho, jak dostat "rozumnou" interpretaci struktury jako string
	fmt.Printf("%#v\n", p)
}
// END,OMIT
