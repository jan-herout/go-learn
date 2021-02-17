package main

import (
"fmt"
"strings"
"sort"
)

// countWords counts all words in the string. Words are split by white space.
func countWords(text string) map[string]int {
	out := make(map[string]int)
	for _,w := range strings.Fields(text) {
		out[w]++
	}
	return out
}

func printSorted(m map[string]int) {
	var keys []string
	for k := range m { keys = append(keys, k) }
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("words[\"%s\"]=%d\n", k, m[k])
	}
}

func main() {
	// START,OMIT
	printSorted(countWords("malý pejsek běžel a běžel setmělým lesem"))
	// END,OMIT
}