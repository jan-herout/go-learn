package main

import (
	"bytes"
	"fmt"
	"io"
)

func dontPanic() {
	if err := recover(); err!=nil {
		fmt.Println("PANIC AVERTED:", err)
	}
}

var bp *bytes.Buffer
var b   bytes.Buffer

// START,OMIT
func read(test string, r io.Reader) {		
	fmt.Println("----------------------------------")
	if _, ok := r.(*bytes.Buffer); ok {
		fmt.Printf("%s: je to pointer na buffer\n%s: adresa je %p\n", test, test, r)
	}
	if r == nil {
		fmt.Printf("%s: je to nil\n", test)
	} else {
		fmt.Printf("%s: NENÍ to nil, můžeme tedy číst?\n", test)
		b := make([]byte,0,10)                // chceme načíst 10 bajtů
		defer dontPanic()                     // defered recover
		n,_ := r.Read(b)	                  // protože tohle ZPANIKAŘÍ, r je NIL interface
		fmt.Println(test, "read",n,"bytes")
	}
}

func main() {
	var r io.Reader                 // nenainicializovaná hodnota
	read("po deklaraci", r)	
	r = bp                          // var bp *bytes.Buffer
	read("r=*bytes.Buffer (nil)", r)
	// proč pointer? Protože jinak: bytes.Buffer does not implement io.Reader (Read method has pointer receiver)
	r = &b                          // var b   bytes.Buffer
	read("r=&bytes.Buffer", r)
}
// END,OMIT
