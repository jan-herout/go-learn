package main

import "fmt"

func main() {
	// uvozovky provádí escape sekvencí, jako je \n, \r, \t, \a
	a := "line 1\nline 2\a"  // \a "rings a bell" pokud to vytisknete na terminál....
	
	// backtick NEPROVÁDÍ interpolaci, a umožňuje roztáhnout literál na více řádek
	// POZOR, Go tvrdošíjně používá jako EOL Linux konvenci, tj pouze LF!
	b := `line 1 \n \a
		line 2`
	
	fmt.Println(a)	
	fmt.Println("-------------------------")
	fmt.Println(b)

}
