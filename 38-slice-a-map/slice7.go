package main

import "fmt"

func main() {
	// START,OMIT		
	var soi []int
	fmt.Printf("%#v\n", soi)
	if soi == nil {
		fmt.Println("a nil slice!!!")
	}
	// soi[0] = 1        // tohle by padlo (panic), protože slice nebyl ještě inicializován
	soi = append(soi, 1) // tohle ale projde
	soi = append(soi, 2)
	soi = append(soi, 3)
	fmt.Printf("%#v\n", soi)
	if soi != nil {
		fmt.Println("*NOT* a nil slice!!!") 
	}
	
	// END,OMIT
}
