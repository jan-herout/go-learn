package main

import "fmt"

func main() {
	sum := 1
	for ; sum<200;  {
		sum += sum
	}
	fmt.Printf("sum=%d\n",sum)
}