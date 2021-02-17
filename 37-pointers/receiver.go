package main

import "fmt"

type posInt int

func (i  posInt) BadAdd(a int) {  i =  i + posInt(a) }
func (i *posInt) Add   (a int) { *i = *i + posInt(a) }

var i posInt

func main() {	
	i.BadAdd(1)               // CHYBA - na kterou si SNADNO naběhnete
	fmt.Println(i)            //         předali jsme kopii čísla 0, a tu povýšili o 1!
	i.Add(1)                  // SPRÁVNĚ - předali jsme *adresu* čísla 0...
	fmt.Println(i)            //         provedli dereferenci, a povýšili o 1
}
