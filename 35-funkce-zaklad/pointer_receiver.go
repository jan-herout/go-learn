package main

import "fmt"
type mujInt int


func (i *mujInt) printVal() { 
    fmt.Println("hodnota je", *i) // *i - dereferencuj i, a vypiš hodnotu na dané adrese
}

//START,OMIT
func (i mujInt) add (what mujInt) mujInt { return i + what } // původní value receiver

// pointer receiver; i je POINTER na typ mujInt
func (i *mujInt) ptrAdd (what mujInt) { 
    // dereferencuj i, a na adresu kam ukazuje i vlož součet hodnot
    *i = *i + what                
}

func main() {    
    var x mujInt
    // vytvoří KOPII objektu x, přidá k ní 10, a vrátí ji; main ale retval zahazuje.
    x.add(10)
    x.printVal()        // = 0
    x.ptrAdd(10)        // x se automaticky převede na pointer, a ten se předá funkci
                        // kompilátor vidí něco jako: (&x).add(10)                        
    x.printVal()        // = 10
}
//END,OMIT