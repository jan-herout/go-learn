package main

import "fmt"

//START,OMIT
type Animal struct { }
type Mammal struct { }
type Dog struct { 
	Animal
	Mammal
}

func (a Animal) Whoami() string { return "?" }	
func (m Mammal) Whoami() string { return "mlask" }

func main() {
	var d Dog
	fmt.Println(d.Whoami())			// prog.go:20:18: ambiguous selector d.Whoami
	fmt.Println(d.Animal.Whoami())	// OK!
	fmt.Println(d.Mammal.Whoami())	// OK!
}
//END,OMIT