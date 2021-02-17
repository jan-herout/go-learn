package main

import (
	"fmt"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func factorial(input string) (int, error) {
	i, err := strconv.Atoi(input)
	if err != nil {
		return -1, fmt.Errorf("can not convert: %w (input was: %s)", err, input)
	}
	if i < 0 {
		return -1, fmt.Errorf("factorial is only defined for positive numbers: %d", i)
	}
	f := 1
	for i > 0 {
		f = f * i
		i--
	}
	return f, nil
}

func main() {
	// START,OMIT
	i, err := factorial("5")
	check(err)
	fmt.Println("factorial(5)=", i)
	i, err = factorial("šest")
	check(err)
	fmt.Println("factorial(šest)=", i)
	// END,OMIT
}
