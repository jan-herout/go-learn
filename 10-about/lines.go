package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
)

var m = regexp.MustCompile(`^\s*(u|f)`)

func main() {
	r, err := os.Open(`c:\BI_Domain\learn-go\10-about\lines.pl`)
	if err != nil {
		panic(err)
	}
	
	scanner := bufio.NewScanner(r)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if m.MatchString(line) {
			fmt.Println(i, line)
		}
		i++
	}
	if scanner.Err() != nil {
		panic(err)
	}
}