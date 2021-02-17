package main

import (
	"fmt"
	"strings"
)

func main() {
	var s strings.Builder
	for i := 0; i < 10; i++ {
		s.WriteString(fmt.Sprintf("hodnota je %d\n", i))
	}
	fmt.Println(s.String())
}
