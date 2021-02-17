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
func (d Dog)    Whoami() string { return "HAF!" }

func main() {
	var d Dog
	fmt.Println(d.Whoami())
	fmt.Println(d.Animal.Whoami())
	fmt.Println(d.Mammal.Whoami())
}
//END,OMIT