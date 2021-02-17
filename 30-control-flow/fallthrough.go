package main

import "fmt"

func main() {
	// START,OMIT
	switch 2 {
	case 1: // false, nespustí se
		fmt.Println("1")
		fallthrough
	case 2: // true, spustí se
		fmt.Println("2") // --- RUN
		fallthrough      // FALLTHROUGH, spustí i následující case
	default:
		fmt.Println("default") // NESPUSTÍ SE (nahoře není fallthrough)
		fallthrough            // **ANO**! I tady je možné ho použít!
	case 3: // podmínka se NEVYHODNOCUJE
		fmt.Println("3") // --- RUN
	}
	// END,OMIT
}
