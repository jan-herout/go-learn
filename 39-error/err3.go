package main

import (
	"fmt"
)

func wontFail() (string, error) { return "all ok!", nil }           // nil = all clear
func willFail() (string, error) { return "", fmt.Errorf("BANG!") }  // vrací error

// START,OMIT
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error
	s, err := wontFail()
	check(err)
	fmt.Println("WE GOT:", s)

	s, err = willFail()
	check(err)
	fmt.Println("WE GOT:", s)
}
// END,OMIT

