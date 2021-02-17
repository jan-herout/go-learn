package main

import (
	"fmt"
	"log"
)

var someInt int

// START,OMIT
func panicky(i int) { _ = 10 / i }

func doWork() (err error) {
	defer func() {                                        // úplně nakonec 
		if pErr := recover(); pErr != nil {               // se podívej, zda nebyl PANIC
			log.Print("recovered from panic: ", pErr)     // pokud ano, zaloguj to
			err = fmt.Errorf("%s", pErr)                  // a vrať tuto informaci upstream
		}
	}()
	panicky(0)     // this *WILL* panic
	return
}

func main() {
	err := doWork()
	if err != nil {
		fmt.Println("doWork returned an error.", err)
	}
	fmt.Println("don't panic!")
}

// END,OMIT
