// START,OMIT
package main

import (
	"fmt"
	"encoding/xml"
)

// Person is ...
type Person struct {
	Name     string `xml:"name"`
	YearBorn int    `xml:"year_born"`
}

// XMLString ignores all errors and returns XML with the struct.
func (p Person) XMLString() string {
	b, _ := xml.Marshal(p)
	return string(b)
}

func main() {
	p := Person{Name: "John Doe", YearBorn: 1977}
	fmt.Println(p.XMLString())
}
// END,OMIT
