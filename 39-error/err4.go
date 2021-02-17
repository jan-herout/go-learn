package main

import (
	"fmt"
	"log"	
)

// START,OMIT
func buggyOpen(path string) error { 
	return fmt.Errorf("failed to open file: %s", path)
}

// analyze returns number of .... or error.
func analyze(name string) (int,error) {
	if err := buggyOpen("/tmp/"+name); err!=nil {
		return 0, fmt.Errorf("analyze: %w", err) // error wrapping - doplnění informací
	}
	return 1, nil // happy path
}

func main() {	
	i, err := analyze("soubor")
	if err!=nil {
		log.Print("číslům nevěř! ", err)
	} else {
		fmt.Println("výsledek je", i)
	}
}
// END,OMIT
