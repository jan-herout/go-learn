package main

import (
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// START,OMIT
	h, err := os.Open(`c:\BI_Domain\learn-go\39-error\io-readall.go`)
	check(err)      // panics on error
	defer h.Close() // close the file on exit

	data, err := io.ReadAll(h) // read WHOLE (!) file into memory
	check(err)
	fmt.Println("file contains", len(data), "bytes")
	// END,OMIT
}
