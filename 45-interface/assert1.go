package main

import (
	"fmt"
)
// START,OMIT
func tryAssert(x interface{}) string {
	if c, ok := x.(fmt.Stringer); ok {
		return fmt.Sprintf("Je to stringer! : %s", c.String())
	}
	if c, ok := x.(error); ok {
		return fmt.Sprintf("Je to chyba! : %s", c.Error())
	}
	if c, ok := x.(int); ok {
		return fmt.Sprintf("Je to int! : %d", c)
	}
	return "Nenám tušení co to je"
}

func main() {
	fmt.Println(tryAssert(100))
	fmt.Println(tryAssert(fmt.Errorf("OOPS")))
	fmt.Println(tryAssert(struct{ name string }{name: "John"}))
}
// END,OMIT