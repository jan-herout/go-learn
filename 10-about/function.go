// TIP: co to podle vás dělá?
package main
import "fmt"
//START,OMIT
func do(a int) { a = a + 10 }
func main() {
	a:=0 
	do(a)
	fmt.Println(a)
}
//END,OMIT