package main

import "fmt"
//START,OMIT
type mujInt int	// nový datový typ, založený na typu int

func (i mujInt) add (what mujInt) mujInt {    // IN: 2x mujInt; OUT: nový mujInt
    return i + what                           // 
}                                             // více až budeme řešit struct{}

func (i mujInt) printVal() {                  // IN: mujInt - value receiver; OUT: nic
    fmt.Println("hodnota je", i)              
}

func main() {
    var x mujInt                     // inicializace, x = 0 (default nulová hodnota)
    x = x.add(5)                     // x = x + 5
    x.add(10).add(20).printVal()     // dočasně vytvoř mujInt jako x.add(10), opakuj s ...
    x.printVal()                     // CO TO VYPÍŠE?
}
//END,OMIT