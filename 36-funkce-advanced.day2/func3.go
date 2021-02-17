package main

import "fmt"

//START,OMIT
type message struct {
	id   int
	text string
}

// generator je factory funkce. vrací generátor zpráv.
func generator(start int) func(string) message {
	id := start                            // id první zprávy je poskytnuté "zvenku"
	return func(text string) message {     // generator je "factory" funkce, vrací funkci, která:
		m := message{id: id, text: text}   // - vytvoří novou message s unikátním ID 
		id++                               // - zvedne ID (CLOSURE, uzavřeno nad vnějším scope)
		return m                           // - vrátí tuto message
	}
}

func main() {
	g1, g1000 := generator(1), generator(1000)
	fmt.Println(g1("první zpráva"))
	fmt.Println(g1("druhá zpráva"))
	fmt.Println(g1000("první zpráva v řadě od tisíce"))
	fmt.Println(g1000("druhá zpráva v řadě od tisíce"))
}
//END,OMIT