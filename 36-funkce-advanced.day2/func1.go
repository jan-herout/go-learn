package main

import "fmt"

// START,OMIT
// funcString je datový typ, který reprezentuje *funkci* s danou signaturou
type funcString func(string)

func printHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	var greet, leave funcString         // dvě proměnné, do kterých jde vložit funkce
	greet = printHello                  // může jít o pojmenovanou funkci (se shodnou signaturou)
	leave = func(name string) {         // ale může jít také o anonymní funkci ...
		fmt.Printf("Bye, %s!\n", name)  //     v lokálním scope
	}

	greet("John")                       // takhle se potom funkce zavolá
	leave("John")

}

// END,OMIT
