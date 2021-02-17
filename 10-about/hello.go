// základní jednotka programu je PACKAGE
// main je SPUSTITELNÁ package
package main	

// import zpřístupňuje balíček "fmt" a balíček "os"
// oba balíčky jsou součástí standardní knihovny

import (
	"fmt"
	"os"
)

// klíčové slovo `func` deklaruje funkci
// funkce `main` je "entry point" programu

func main() {
	// exportované funkce a/nebo proměnné z jiných balíčků
	// jsou dostupné pomocí "tečkové" konvence
	fmt.Println("Hello world")
	
	// pokud *opravdu* potřebujete vynutit návratovou hodnotu z programu...
	os.Exit(1)
}
