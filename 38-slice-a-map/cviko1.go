package main

import (
"fmt"
"strings"
)

func firstRune(s string) string {
	for _,r := range s { return strings.ToLower(string(r)) }
	return ""
}

func filterVowels(line string) []string {
	in := strings.Fields(line)
	out := make([]string,0,len(in))
	for _, v := range in {
		if ! strings.ContainsAny(firstRune(v),"aeiouy") {
			out = append(out, v)
		}
	}
	return out
}
func main() {
	// START,OMIT
	fmt.Printf("%#v\n", filterVowels("ementál není smrdutý sýr, ale pěkně voní"))
	// END,OMIT
}