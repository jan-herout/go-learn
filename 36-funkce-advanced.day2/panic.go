package main

import (
	"fmt"
)

var someInt int

// START,OMIT
func panicky(s string) { panic(s) }

func doWork() {
	panicky("BANG!")     // this *WILL* panic
}

func main() {
	doWork()
	fmt.Println("don't panic!")
}

// END,OMIT
