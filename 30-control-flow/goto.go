package main

import "fmt"

// START,OMIT
func main() {
	fmt.Println("start")
	goto End
	fmt.Println("skip")
End:
	fmt.Println("end")
}
// END,OMIT
