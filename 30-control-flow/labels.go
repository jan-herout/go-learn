package main

import "fmt"

func main(){
	// START,OMIT
	// Outer: je "štítek" (label), který označuje vnější smyčku 
	Outer: for o:=1; o<10; o++ {
		
		for i:=5; i<10; i++ {
			fmt.Println(o,i)
			if o > i {
				fmt.Println("break!")
				
				// Break/continue bez parametrů ukončuje *vnitřní* cyklus.
				// Můžete ale zadat název štítku, který chcete přerušit.
				break Outer
			}
		}
	}
	// END,OMIT
}